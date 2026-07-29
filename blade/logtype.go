package blade

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed logtyped.yml
var logTypedYML []byte

// fieldConfig describes one output column:
//   - key:        JSON field name, supports dot-notation for nested fields (e.g. body.log)
//   - label:      column header printed before the value
//   - width:      fixed column width; 0 = no constraint
//   - transform:  optional value transform (json_text_array, single_line)
//   - value_map:  map raw value → display string (e.g. "false" → "Message")
type fieldConfig struct {
	Key       string            `yaml:"key"`
	Label     string            `yaml:"label"`
	Width     int               `yaml:"width"`
	Transform string            `yaml:"transform,omitempty"`
	ValueMap  map[string]string `yaml:"value_map,omitempty"`
}

// eventRule selects a field set when the event_type_key value equals Match.
type eventRule struct {
	Match  string        `yaml:"match"`
	Fields []fieldConfig `yaml:"fields"`
}

// logTypeEntry is one named log-type definition from logtyped.yml.
//
// Filtering:
//   - filter:    key→value pairs; ALL must match (exact string equality).
//   - filter_in: key→[]values; at least one value must match per key.
//
// Field selection:
//   - Simple mode (fields):      same field list for every event.
//   - Conditional mode (event_type_key + event_rules): field list chosen by
//     matching the value of event_type_key; falls back to default_fields.
//
// Non-JSON messages and filtered-out events are skipped silently; the
// sequence counter does not advance for them.
type logTypeEntry struct {
	ID          int    `yaml:"id"`
	Description string `yaml:"description"`

	Filter   map[string]string   `yaml:"filter,omitempty"`
	FilterIn map[string][]string `yaml:"filter_in,omitempty"`

	Fields []fieldConfig `yaml:"fields,omitempty"`

	EventTypeKey  string        `yaml:"event_type_key,omitempty"`
	DefaultFields []fieldConfig `yaml:"default_fields,omitempty"`
	EventRules    []eventRule   `yaml:"event_rules,omitempty"`
}

type logTypesFile struct {
	LogTypes []logTypeEntry `yaml:"logtypes"`
}

// logTypeFormatter formats a parsed JSON log message with a sequence number.
// An empty return value signals the caller to skip this event.
type logTypeFormatter func(seq int, data map[string]interface{}) string

// logTypeFormatters is populated at init time from the embedded logtyped.yml.
var logTypeFormatters = map[int]logTypeFormatter{}

func init() {
	var cfg logTypesFile
	if err := yaml.Unmarshal(logTypedYML, &cfg); err != nil {
		panic(fmt.Errorf("parse logtyped.yml: %w", err))
	}
	for _, lt := range cfg.LogTypes {
		lt := lt
		logTypeFormatters[lt.ID] = buildFormatter(lt)
	}
}

func buildFormatter(lt logTypeEntry) logTypeFormatter {
	return func(seq int, data map[string]interface{}) string {
		// Exact-match filters: all conditions must hold.
		for key, required := range lt.Filter {
			if extractField(data, key) != required {
				return ""
			}
		}
		// Any-of filters: at least one allowed value must match per key.
		for key, allowed := range lt.FilterIn {
			actual := extractField(data, key)
			found := false
			for _, v := range allowed {
				if actual == v {
					found = true
					break
				}
			}
			if !found {
				return ""
			}
		}

		// Pick the right field list for this event.
		fields := lt.Fields
		if lt.EventTypeKey != "" {
			eventTypeVal := extractField(data, lt.EventTypeKey)
			fields = lt.DefaultFields
			for _, rule := range lt.EventRules {
				if rule.Match == eventTypeVal {
					fields = rule.Fields
					break
				}
			}
		}

		var sb strings.Builder
		fmt.Fprintf(&sb, "[%04d]", seq)
		for _, f := range fields {
			val := extractField(data, f.Key)
			if len(f.ValueMap) > 0 {
				if mapped, ok := f.ValueMap[val]; ok {
					val = mapped
				}
			}
			val = applyTransform(val, f.Transform)
			val = strings.TrimSpace(val)
			fmt.Fprintf(&sb, " %s: %s", f.Label, fixedWidth(val, f.Width))
		}
		return sb.String()
	}
}

// extractField retrieves a (possibly nested) field from data using dot notation.
// e.g. "body.log" reads data["body"]["log"].
func extractField(data map[string]interface{}, key string) string {
	i := strings.IndexByte(key, '.')
	if i < 0 {
		return anyToString(data[key])
	}
	nested, _ := data[key[:i]].(map[string]interface{})
	if nested == nil {
		return ""
	}
	return extractField(nested, key[i+1:])
}

// anyToString converts a JSON-unmarshalled value to its string representation.
func anyToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", val)
	}
}

// applyTransform applies the named transform to a raw field value.
//
//   - json_text_array: val is a JSON-encoded array of {type,value} objects;
//     extract and join the "text" values.
//   - single_line: replace newline characters with spaces.
func applyTransform(val, transform string) string {
	switch transform {
	case "json_text_array":
		var items []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		}
		if err := json.Unmarshal([]byte(val), &items); err != nil {
			return val
		}
		var parts []string
		for _, item := range items {
			if item.Type == "text" && item.Value != "" {
				parts = append(parts, strings.TrimSpace(item.Value))
			}
		}
		return strings.Join(parts, " ")
	case "single_line":
		val = strings.ReplaceAll(val, "\r\n", " ")
		val = strings.ReplaceAll(val, "\n", " ")
		return val
	}
	return val
}

// fixedWidth pads s to width with spaces (left-aligned), or truncates if longer.
// width 0 means no constraint: the value is returned as-is.
func fixedWidth(s string, width int) string {
	if width == 0 {
		return s
	}
	if len(s) > width {
		return s[:width]
	}
	return fmt.Sprintf("%-*s", width, s)
}

// formatWithLogType applies the logtype formatter to a raw CloudWatch log
// message. Returns (formatted, true) on success, ("", true) when filtered out,
// or (rawMessage, false) when the logtype is unknown or the message is not JSON.
func formatWithLogType(logType, seq int, rawMessage string) (string, bool) {
	formatter, ok := logTypeFormatters[logType]
	if !ok {
		return rawMessage, false
	}

	msg := strings.TrimSpace(rawMessage)
	data := map[string]interface{}{}
	if err := json.Unmarshal([]byte(msg), &data); err != nil {
		return rawMessage, false
	}

	return formatter(seq, data), true
}
