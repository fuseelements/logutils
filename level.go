package logutils

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

type LogLevel string

// LevelFilter is an io.Writer that can be used with a logger that
// will filter out log messages that aren't at least a certain level.
//
// Once the filter is in use somewhere, it is not safe to modify
// the structure.
type LevelFilter struct {
	// Levels is the list of log levels, in increasing order of
	// severity. Example might be: {"DEBUG", "WARN", "ERROR"}.
	Levels []LogLevel

	// MinLevel is the minimum level allowed through
	MinLevel LogLevel

	// The underlying io.Writer where log messages that pass the filter
	// will be set.
	Writer io.Writer

	badLevels map[LogLevel]struct{}
	once      sync.Once

	// Enable coloured output based on level
	Color bool
}

var colors map[LogLevel]string

type color int

const (
	colorBlack = (iota + 30)
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite
)

// Check will check a given line if it would be included in the level
// filter.
func (f *LevelFilter) Check(line []byte) bool {
	f.once.Do(f.init)

	// Check for a log level
	var level LogLevel = getLevel(line)

	_, ok := f.badLevels[level]
	return !ok
}

func (f *LevelFilter) Write(p []byte) (n int, err error) {
	// Note in general that io.Writer can receive any byte sequence
	// to write, but the "log" package always guarantees that we only
	// get a single line. We use that as a slight optimization within
	// this method, assuming we're dealing with a single, complete line
	// of log data.

	if !f.Check(p) {
		return len(p), nil
	}

	if f.Color == true {
		var level LogLevel = getLevel(p)
		buf := &bytes.Buffer{}

		if colorStart, ok := colors[level]; ok {
			buf.Write([]byte(colorStart))
			buf.Write(p)
			buf.Write([]byte(closeColorSeq()))
		} else {
			buf.Write(p)
		}

		fmt.Println(buf.String())

		return f.Writer.Write(buf.Bytes())
	}

	return f.Writer.Write(p)
}

func (f *LevelFilter) init() {
	f.badLevels = make(map[LogLevel]struct{})
	for _, level := range f.Levels {
		if level == f.MinLevel {
			break
		}

		f.badLevels[level] = struct{}{}
	}
}

func colorSeq(color color) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

func closeColorSeq() string {
	return "\033[0m"
}

func init() {
	colors = map[LogLevel]string{
		"CRITICAL": colorSeq(colorMagenta),
		"CRIT":     colorSeq(colorMagenta),
		"ERROR":    colorSeq(colorRed),
		"WARNING":  colorSeq(colorYellow),
		"WARN":     colorSeq(colorYellow),
		"NOTICE":   colorSeq(colorGreen),
		"DEBUG":    colorSeq(colorCyan),
	}
}

func getLevel(line []byte) (level LogLevel) {
	x := bytes.IndexByte(line, '[')
	if x >= 0 {
		y := bytes.IndexByte(line[x:], ']')
		if y >= 0 {
			level = LogLevel(line[x+1 : x+y])
		}
	}
	return
}
