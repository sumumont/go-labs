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
	"github.com/rs/zerolog"
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
	zerolog.TimeFieldFormat = timeFormat

	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
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
	return logger.Err(err).Caller().Str(goroutineId, GetGoroutineId())
}

func Fatal() *zerolog.Event {
	return logger.Fatal().Caller()
}

//后续接入了链路追踪，可以从ctx里面取出来
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
