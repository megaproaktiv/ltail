package blade

import (
	"fmt"
	"strings"
	"testing"
)

func TestFormatWithLogType(t *testing.T) {
	msg1 := `{ "ContactFlowModuleType": "PlayPrompt", "ContactFlowName": "Default customer whisper", "ContactId": "470d9779" }`
	msg2 := `{ "ContactFlowModuleType": "Resume", "ContactFlowName": "Default customer whisper", "Parameters": {} }`

	cases := []struct {
		seq  int
		msg  string
		typ  string
		name string
	}{
		{1, msg1, "PlayPrompt", "Default customer whisper"},
		{2, msg2, "Resume", "Default customer whisper"},
	}

	for _, c := range cases {
		got, ok := formatWithLogType(1, c.seq, c.msg)
		if !ok {
			t.Errorf("seq %d: formatWithLogType returned ok=false", c.seq)
			continue
		}
		wantPrefix := fmt.Sprintf("[%04d]", c.seq)
		if !strings.HasPrefix(got, wantPrefix) {
			t.Errorf("seq %d: missing sequence prefix, got %q", c.seq, got)
		}
		// Type column: fixed 64 chars
		wantType := fmt.Sprintf("%-64s", c.typ)
		if !strings.Contains(got, "Type: "+wantType) {
			t.Errorf("seq %d: Type column not fixed-width 64\n got: %q", c.seq, got)
		}
		// Name column: fixed 32 chars
		wantName := fmt.Sprintf("%-32s", c.name)
		if !strings.Contains(got, "Name: "+wantName) {
			t.Errorf("seq %d: Name column not fixed-width 32\n got: %q", c.seq, got)
		}
	}

	// Non-JSON falls back gracefully
	_, ok := formatWithLogType(1, 3, "not json at all")
	if ok {
		t.Error("expected ok=false for non-JSON message")
	}

	// Unknown logtype returns false
	_, ok = formatWithLogType(99, 1, msg1)
	if ok {
		t.Error("expected ok=false for unknown logtype")
	}
}

func TestFixedWidth(t *testing.T) {
	cases := []struct {
		input string
		width int
		want  string
	}{
		{"hello", 10, "hello     "},
		{"hello", 5, "hello"},
		{"toolongvalue", 5, "toolo"},
		{"", 4, "    "},
	}
	for _, c := range cases {
		got := fixedWidth(c.input, c.width)
		if got != c.want {
			t.Errorf("fixedWidth(%q, %d) = %q, want %q", c.input, c.width, got, c.want)
		}
		if len(got) != c.width {
			t.Errorf("fixedWidth(%q, %d): len=%d, want %d", c.input, c.width, len(got), c.width)
		}
	}
}
