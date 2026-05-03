package log_parser

import (
	"io"
	"time"
)

// Level is a log severity level.
type Level string

const (
	DEBUG Level = "DEBUG"
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
)

// Entry is a successfully parsed log line.
type Entry struct {
	Timestamp time.Time
	Level     Level
	Service   string
	Message   string
}

// ParseError records a line that could not be parsed.
type ParseError struct {
	Line int
	Raw  string
	Err  error
}

// Result holds the outcome of parsing a full log stream.
type Result struct {
	Entries []Entry
	Errors  []ParseError
}

// ParseLog reads log lines from r. Valid lines are parsed into Entries;
// malformed lines are recorded in Errors. Parsing never stops early.
//
// Expected line format:
//
//	2006-01-02T15:04:05Z INFO service-name: message text here
//
// Hint: use bufio.Scanner to read line by line; strings.SplitN to split fields.
func ParseLog(r io.Reader) Result {
	// TODO: implement
	return Result{}
}

// FilterByLevel returns entries at exactly the given level.
func FilterByLevel(entries []Entry, level Level) []Entry {
	// TODO: implement
	return nil
}

// Between returns entries whose timestamp falls within [start, end] inclusive.
func Between(entries []Entry, start, end time.Time) []Entry {
	// TODO: implement
	return nil
}
