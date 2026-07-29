package blade

import (
	"fmt"
	"strings"
	"testing"
)

// ── logtype 1 ────────────────────────────────────────────────────────────────

func TestFormatWithLogType1(t *testing.T) {
	msg1 := `{ "ContactFlowModuleType": "PlayPrompt", "ContactFlowName": "Default customer whisper" }`
	msg2 := `{ "ContactFlowModuleType": "Resume",     "ContactFlowName": "Default customer whisper" }`

	cases := []struct{ seq int; msg, typ, name string }{
		{1, msg1, "PlayPrompt", "Default customer whisper"},
		{2, msg2, "Resume", "Default customer whisper"},
	}

	for _, c := range cases {
		got, ok := formatWithLogType(1, c.seq, c.msg)
		if !ok {
			t.Fatalf("seq %d: ok=false", c.seq)
		}
		if !strings.HasPrefix(got, fmt.Sprintf("[%04d]", c.seq)) {
			t.Errorf("seq %d: missing sequence prefix in %q", c.seq, got)
		}
		if want := "Type: " + fmt.Sprintf("%-64s", c.typ); !strings.Contains(got, want) {
			t.Errorf("seq %d: Type column wrong\n got:  %q\n want substring: %q", c.seq, got, want)
		}
		if want := "Name: " + fmt.Sprintf("%-32s", c.name); !strings.Contains(got, want) {
			t.Errorf("seq %d: Name column wrong\n got:  %q\n want substring: %q", c.seq, got, want)
		}
	}
}

// ── logtype 2 ────────────────────────────────────────────────────────────────

func TestFormatWithLogType2_CreateSession(t *testing.T) {
	msg := `{ "event_type": "TRANSCRIPT_CREATE_SESSION", "session_id": "abc" }`
	got, ok := formatWithLogType(2, 1, msg)
	if !ok {
		t.Fatalf("ok=false")
	}
	want := "[0001] Type: TRANSCRIPT_CREATE_SESSION"
	if got != want {
		t.Errorf("\n got:  %q\n want: %q", got, want)
	}
}

func TestFormatWithLogType2_OrchestrationMessage(t *testing.T) {
	msg := `{
		"event_type": "TRANSCRIPT_ORCHESTRATION_MESSAGE",
		"participant": "CUSTOMER",
		"values": "[{\"type\":\"text\",\"value\":\"ich habe ein problem\"}]"
	}`
	got, ok := formatWithLogType(2, 2, msg)
	if !ok {
		t.Fatalf("ok=false")
	}
	if !strings.Contains(got, "[0002]") {
		t.Errorf("missing sequence: %q", got)
	}
	if !strings.Contains(got, "Type: TRANSCRIPT_ORCHESTRATION_MESSAGE") {
		t.Errorf("missing Type: %q", got)
	}
	if !strings.Contains(got, "Party: CUSTOMER") {
		t.Errorf("missing Party: %q", got)
	}
	if !strings.Contains(got, "Message: ich habe ein problem") {
		t.Errorf("missing Message: %q", got)
	}
}

func TestFormatWithLogType2_AgenticMessage(t *testing.T) {
	msg := `{
		"event_type": "TRANSCRIPT_AGENTIC_MESSAGE",
		"parsed_response": "\nVerstehe, Sie haben eine Frage.\n"
	}`
	got, ok := formatWithLogType(2, 4, msg)
	if !ok {
		t.Fatalf("ok=false")
	}
	if !strings.Contains(got, "Type: TRANSCRIPT_AGENTIC_MESSAGE") {
		t.Errorf("missing Type: %q", got)
	}
	if !strings.Contains(got, "parsed_response: Verstehe, Sie haben eine Frage.") {
		t.Errorf("missing parsed_response (trimmed): %q", got)
	}
}

func TestFormatWithLogType2_UnknownEventType(t *testing.T) {
	msg := `{ "event_type": "SOME_OTHER_EVENT", "foo": "bar" }`
	got, ok := formatWithLogType(2, 5, msg)
	if !ok {
		t.Fatalf("ok=false")
	}
	// Falls back to default_fields: only Type
	if !strings.Contains(got, "Type: SOME_OTHER_EVENT") {
		t.Errorf("expected default Type field: %q", got)
	}
	if strings.Contains(got, "Party") || strings.Contains(got, "Message") {
		t.Errorf("unexpected fields in default fallback: %q", got)
	}
}

func TestFormatWithLogType3_InfoShown(t *testing.T) {
	msg := `{"level":"INFO","msg":"handler invoked","time":"2026-07-29T15:40:27Z"}`
	got, ok := formatWithLogType(3, 1, msg)
	if !ok {
		t.Fatalf("ok=false")
	}
	want := "[0001] Level: INFO Message: handler invoked"
	if got != want {
		t.Errorf("\n got:  %q\n want: %q", got, want)
	}
}

func TestFormatWithLogType3_DebugSkipped(t *testing.T) {
	msg := `{"level":"DEBUG","msg":"Calling Parameter Store","time":"2026-07-29T15:40:27Z"}`
	got, ok := formatWithLogType(3, 1, msg)
	if !ok {
		t.Fatalf("ok=false")
	}
	if got != "" {
		t.Errorf("expected empty string (filtered), got %q", got)
	}
}

func TestFormatWithLogType3_NonJSONSkipped(t *testing.T) {
	// Lambda START/END/REPORT lines are not JSON — must return ok=false.
	lines := []string{
		"START RequestId: bde95ec5 Version: $LATEST",
		"END RequestId: bde95ec5",
		"REPORT RequestId: bde95ec5 Duration: 353.74 ms",
	}
	for _, line := range lines {
		_, ok := formatWithLogType(3, 1, line)
		if ok {
			t.Errorf("expected ok=false for non-JSON line %q", line)
		}
	}
}

// ── logtype 4 ────────────────────────────────────────────────────────────────

func TestFormatWithLogType4_InfoMessage(t *testing.T) {
	msg := `{
		"severityText": "INFO",
		"body": { "isError": false, "log": "Executing tool get-available-callback-slots" }
	}`
	got, ok := formatWithLogType(4, 1, msg)
	if !ok {
		t.Fatalf("ok=false")
	}
	want := "[0001] Type: Message Log: Executing tool get-available-callback-slots"
	if got != want {
		t.Errorf("\n got:  %q\n want: %q", got, want)
	}
}

func TestFormatWithLogType4_InfoError(t *testing.T) {
	msg := `{
		"severityText": "INFO",
		"body": { "isError": true, "log": "something went wrong" }
	}`
	got, ok := formatWithLogType(4, 2, msg)
	if !ok {
		t.Fatalf("ok=false")
	}
	if !strings.Contains(got, "Type: Error") {
		t.Errorf("expected Type: Error, got %q", got)
	}
	if !strings.Contains(got, "Log: something went wrong") {
		t.Errorf("expected Log field, got %q", got)
	}
}

func TestFormatWithLogType4_WarnShown(t *testing.T) {
	msg := `{ "severityText": "WARN", "body": { "isError": false, "log": "warn msg" } }`
	got, ok := formatWithLogType(4, 3, msg)
	if !ok || got == "" {
		t.Errorf("WARN should be shown, got %q ok=%v", got, ok)
	}
}

func TestFormatWithLogType4_DebugSkipped(t *testing.T) {
	msg := `{ "severityText": "DEBUG", "body": { "isError": false, "log": "debug msg" } }`
	got, ok := formatWithLogType(4, 1, msg)
	if !ok {
		t.Fatalf("ok=false")
	}
	if got != "" {
		t.Errorf("DEBUG should be filtered, got %q", got)
	}
}

// ── extractField / anyToString ────────────────────────────────────────────────

func TestExtractField(t *testing.T) {
	data := map[string]interface{}{
		"top": "value",
		"body": map[string]interface{}{
			"log":     "hello",
			"isError": false,
			"deep": map[string]interface{}{
				"val": "nested",
			},
		},
	}
	cases := []struct{ key, want string }{
		{"top", "value"},
		{"body.log", "hello"},
		{"body.isError", "false"},
		{"body.deep.val", "nested"},
		{"missing", ""},
		{"body.missing", ""},
	}
	for _, c := range cases {
		got := extractField(data, c.key)
		if got != c.want {
			t.Errorf("extractField(%q) = %q, want %q", c.key, got, c.want)
		}
	}
}

// ── helpers ───────────────────────────────────────────────────────────────────

func TestFixedWidth(t *testing.T) {
	cases := []struct {
		input string
		width int
		want  string
	}{
		{"hello", 10, "hello     "},
		{"hello", 5, "hello"},
		{"toolong", 5, "toolo"},
		{"", 4, "    "},
		{"anything", 0, "anything"}, // width 0 = no constraint
	}
	for _, c := range cases {
		got := fixedWidth(c.input, c.width)
		if got != c.want {
			t.Errorf("fixedWidth(%q, %d) = %q, want %q", c.input, c.width, got, c.want)
		}
		if c.width > 0 && len(got) != c.width {
			t.Errorf("fixedWidth(%q, %d): len=%d, want %d", c.input, c.width, len(got), c.width)
		}
	}
}

func TestApplyTransform(t *testing.T) {
	t.Run("json_text_array", func(t *testing.T) {
		in := `[{"type":"text","value":"hello world"}]`
		got := applyTransform(in, "json_text_array")
		if got != "hello world" {
			t.Errorf("got %q", got)
		}
	})
	t.Run("json_text_array with newline", func(t *testing.T) {
		in := `[{"type":"text","value":"\nhello\n"}]`
		got := applyTransform(in, "json_text_array")
		if got != "hello" {
			t.Errorf("got %q", got)
		}
	})
	t.Run("single_line", func(t *testing.T) {
		got := applyTransform("\nhello\nworld\n", "single_line")
		if got != " hello world " {
			t.Errorf("got %q", got)
		}
	})
	t.Run("unknown transform is noop", func(t *testing.T) {
		got := applyTransform("val", "nonexistent")
		if got != "val" {
			t.Errorf("got %q", got)
		}
	})
}
