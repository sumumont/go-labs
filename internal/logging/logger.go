///////////////////////////////////////////////////////////////////////////////
//
// Copyright (c) 2019-present Apulis Technology (Shenzhen) Incorporated. All Rights Reserved
//
//
// Distributed under the MIT License (http://opensource.org/licenses/MIT)
//
///////////////////////////////////////////////////////////////////////////////

package logging

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
	"runtime"
	"strconv"
)

const timeFormat = "2006-01-02 15:04:05"
const goroutineId = "goroutineId"

var logger zerolog.Logger

func init() {
	zerolog.CallerSkipFrameCount = 3
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger = zerolog.New(os.Stdout).With().Logger()
	//zerolog.TimeFieldFormat = timeFormat
	//output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
	//output.FormatLevel = func(i interface{}) string {
	//	return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	//}
	//output.FormatCaller = func(i interface{}) string {
	//	return fmt.Sprintf("%s", i)
	//}
	//{
	//	output.FormatFieldName = func(i interface{}) string {
	//		return fmt.Sprintf("%s=", i)
	//	}
	//	output.FormatFieldValue = func(i interface{}) string {
	//		return fmt.Sprintf("%s", i)
	//	}
	//
	//	output.FormatErrFieldName = func(i interface{}) string {
	//		return fmt.Sprintf("%s=", i)
	//	}
	//	output.FormatErrFieldValue = func(i interface{}) string {
	//		return fmt.Sprintf("%s", i)
	//	}
	//}
	//
	//output.FormatMessage = func(i interface{}) string {
	//	if i == nil {
	//		return fmt.Sprintf("Msg=")
	//	}
	//	return fmt.Sprintf("Msg=%s", i)
	//}
	//output.PartsOrder = []string{zerolog.TimestampFieldName, zerolog.LevelFieldName, "goroutineId", zerolog.CallerFieldName, zerolog.MessageFieldName}
	//output.PartsExclude = []string{"goroutineId"}
	//logger = zerolog.New(output).With().Timestamp().Logger()
}

func SetOutput(w io.Writer) zerolog.Logger {
	logger = zerolog.New(w).With().Timestamp().Logger()
	return logger
}

func Debug() *zerolog.Event {
	return logger.Debug().Caller().Str(goroutineId, GetGoroutineId())
}

func Info() *zerolog.Event {
	return logger.Info().Caller().Str(goroutineId, GetGoroutineId())
}

func Warn() *zerolog.Event {
	return logger.Warn().Caller().Str(goroutineId, GetGoroutineId())
}

func Error(err error) *zerolog.Event {
	return logger.Error().Caller().Str(goroutineId, GetGoroutineId()).Stack().Err(errors.New(err.Error()))
}

func Fatal() *zerolog.Event {
	return logger.Fatal().Caller().Str(goroutineId, GetGoroutineId())
}

// 后续接入了链路追踪，可以从ctx里面取出来
func GetGoroutineId() string {
	goroutineId := strconv.FormatUint(GetGID(), 10)
	return goroutineId
}
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
