package log_parser

import (
	"strings"
	"testing"
	"time"
)

const validLog = `2024-01-15T10:30:00Z INFO api: request received
2024-01-15T10:30:01Z ERROR auth: login failed for user alice
2024-01-15T10:30:02Z WARN db: slow query detected
2024-01-15T10:30:03Z DEBUG cache: cache miss for key session-42`

func TestParseLog_ValidLines(t *testing.T) {
	result := ParseLog(strings.NewReader(validLog))

	if len(result.Errors) != 0 {
		t.Errorf("unexpected parse errors: %v", result.Errors)
	}
	if len(result.Entries) != 4 {
		t.Fatalf("got %d entries, want 4", len(result.Entries))
	}

	e := result.Entries[1]
	if e.Level != ERROR {
		t.Errorf("entry[1].Level = %q, want ERROR", e.Level)
	}
	if e.Service != "auth" {
		t.Errorf("entry[1].Service = %q, want \"auth\"", e.Service)
	}
	if e.Message != "login failed for user alice" {
		t.Errorf("entry[1].Message = %q, want \"login failed for user alice\"", e.Message)
	}
}

func TestParseLog_MalformedLinesCollected(t *testing.T) {
	input := `2024-01-15T10:30:00Z INFO api: ok
this is not a log line
also bad
2024-01-15T10:30:02Z DEBUG svc: done`

	result := ParseLog(strings.NewReader(input))

	if len(result.Entries) != 2 {
		t.Errorf("got %d entries, want 2", len(result.Entries))
	}
	if len(result.Errors) != 2 {
		t.Errorf("got %d parse errors, want 2", len(result.Errors))
	}
	if result.Errors[0].Line != 2 {
		t.Errorf("first parse error on line %d, want 2", result.Errors[0].Line)
	}
	if result.Errors[0].Raw == "" {
		t.Error("ParseError.Raw should contain the original line")
	}
}

func TestParseLog_InvalidTimestamp(t *testing.T) {
	input := "not-a-timestamp INFO svc: msg\n2024-01-15T10:30:00Z INFO svc: ok\n"
	result := ParseLog(strings.NewReader(input))
	if len(result.Errors) != 1 {
		t.Errorf("got %d errors, want 1", len(result.Errors))
	}
	if len(result.Entries) != 1 {
		t.Errorf("got %d entries, want 1", len(result.Entries))
	}
}

func TestParseLog_EmptyInput(t *testing.T) {
	result := ParseLog(strings.NewReader(""))
	if len(result.Entries) != 0 || len(result.Errors) != 0 {
		t.Error("expected empty result for empty input")
	}
}

func TestFilterByLevel(t *testing.T) {
	result := ParseLog(strings.NewReader(validLog))
	errors := FilterByLevel(result.Entries, ERROR)
	if len(errors) != 1 {
		t.Errorf("got %d ERROR entries, want 1", len(errors))
	}
	if errors[0].Service != "auth" {
		t.Errorf("wrong entry returned: %+v", errors[0])
	}

	debugs := FilterByLevel(result.Entries, DEBUG)
	if len(debugs) != 1 {
		t.Errorf("got %d DEBUG entries, want 1", len(debugs))
	}
}

func TestBetween_InclusiveBounds(t *testing.T) {
	result := ParseLog(strings.NewReader(validLog))

	start, _ := time.Parse(time.RFC3339, "2024-01-15T10:30:01Z")
	end, _ := time.Parse(time.RFC3339, "2024-01-15T10:30:02Z")

	got := Between(result.Entries, start, end)
	if len(got) != 2 {
		t.Errorf("got %d entries in range, want 2", len(got))
	}
}

func TestBetween_ExcludesOutside(t *testing.T) {
	result := ParseLog(strings.NewReader(validLog))

	start, _ := time.Parse(time.RFC3339, "2024-01-15T10:30:05Z")
	end, _ := time.Parse(time.RFC3339, "2024-01-15T10:30:10Z")

	got := Between(result.Entries, start, end)
	if len(got) != 0 {
		t.Errorf("expected no entries in future range, got %d", len(got))
	}
}
