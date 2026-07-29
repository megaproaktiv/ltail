package blade

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed logtyped.yml
var logTypedYML []byte

// fieldConfig defines one extracted field: the JSON key, display label, and column width.
type fieldConfig struct {
	Key   string `yaml:"key"`
	Label string `yaml:"label"`
	Width int    `yaml:"width"`
}

// logTypeEntry is one named log-type definition from logtyped.yml.
type logTypeEntry struct {
	ID          int           `yaml:"id"`
	Description string        `yaml:"description"`
	Fields      []fieldConfig `yaml:"fields"`
}

type logTypesFile struct {
	LogTypes []logTypeEntry `yaml:"logtypes"`
}

// logTypeFormatter formats a parsed JSON log message with a sequence number.
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
		logTypeFormatters[lt.ID] = func(seq int, data map[string]interface{}) string {
			var sb strings.Builder
			fmt.Fprintf(&sb, "[%04d]", seq)
			for _, f := range lt.Fields {
				val, _ := data[f.Key].(string)
				fmt.Fprintf(&sb, " %s: %s", f.Label, fixedWidth(val, f.Width))
			}
			return sb.String()
		}
	}
}

// fixedWidth pads s to width with spaces (left-aligned), or truncates if longer.
func fixedWidth(s string, width int) string {
	if len(s) > width {
		return s[:width]
	}
	return fmt.Sprintf("%-*s", width, s)
}

// formatWithLogType applies the logtype formatter to a raw CloudWatch log
// message. It returns the formatted line and true on success, or the original
// message and false when the logtype is unknown or the message is not JSON.
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
