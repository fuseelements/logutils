package logutils

import (
	"bytes"
	"io"
	"log"
	"testing"
)

func TestLevelFilter_impl(t *testing.T) {
	var _ io.Writer = new(LevelFilter)
}

func TestLevelFilter(t *testing.T) {
	buf := new(bytes.Buffer)
	filter := &LevelFilter{
		Levels:   []LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: "WARN",
		Writer:   buf,
	}

	logger := log.New(filter, "", 0)
	logger.Print("[WARN] foo")
	logger.Println("[ERROR] bar")
	logger.Println("[DEBUG] baz")
	logger.Println("[WARN] buzz")

	result := buf.String()
	expected := "[WARN] foo\n[ERROR] bar\n[WARN] buzz\n"
	if result != expected {
		t.Fatalf("bad: %#v", result)
	}
}

func TestLevelFilterCheck(t *testing.T) {
	filter := &LevelFilter{
		Levels:   []LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: "WARN",
		Writer:   nil,
	}

	testCases := []struct {
		line  string
		check bool
	}{
		{"[WARN] foo\n", true},
		{"[ERROR] bar\n", true},
		{"[DEBUG] baz\n", false},
		{"[WARN] buzz\n", true},
	}

	for _, testCase := range testCases {
		result := filter.Check([]byte(testCase.line))
		if result != testCase.check {
			t.Errorf("Fail: %s", testCase.line)
		}
	}
}

func TestColorFormatting(t *testing.T) {
	buf := new(bytes.Buffer)

	filter := NewFilter(buf, true)

	testCases := []struct {
		line     string
		expected string
	}{
		{"[WARN] foo\n", "\x1b[33m[WARN] foo\n\x1b[0m"},
		{"[ERROR] bar\n", "\x1b[31m[ERROR] bar\n\x1b[0m"},
		{"[DEBUG] baz\n", "\x1b[36m[DEBUG] baz\n\x1b[0m"},
		{"[WARN] buzz\n", "\x1b[33m[WARN] buzz\n\x1b[0m"},
	}

	logger := log.New(filter, "", 0)

	for _, testCase := range testCases {
		logger.Print(testCase.line)

		result := buf.String()
		if result != testCase.expected {
			t.Errorf("bad: %#v", result)
		}

		buf.Reset()
	}
}
