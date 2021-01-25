package log

import (
	"bytes"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var buffer bytes.Buffer

func init() {
	//use UNIX Time since it's faster and smaller
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// pretty outputfor console output
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

// D log debug message
func D(m ...interface{}) {
	zlog.WithLevel(zerolog.DebugLevel).Msg(toString(m...))
}

// E log error message
func E(m ...interface{}) {
	zlog.WithLevel(zerolog.ErrorLevel).Msg(toString(m...))
}

// F log fatal message
func F(m ...interface{}) {
	zlog.WithLevel(zerolog.FatalLevel).Msg(toString(m...))
}

// I log info message
func I(m ...interface{}) {
	zlog.WithLevel(zerolog.InfoLevel).Msg(toString(m...))
}

// W log warning message
func W(m ...interface{}) {
	zlog.WithLevel(zerolog.WarnLevel).Msg(toString(m...))
}

func toString(args ...interface{}) string {
	buffer.Reset()
	size := len(args) - 1
	for i, arg := range args {
		buffer.WriteString(fmt.Sprint(arg))
		if i < size {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}
