/*
* (C) 2001-2015 gtlService Inc.
*
* gutils source code
* version: 1.0.0
* author: bobliu0909@gmail.com
* datetime: 2015-12-20
*
 */

package logger

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const DEFAULT_MAXLOGSIZE int64 = 20971520

var glogger *Glogger

type Args struct {
	FileName string
	Level    string
	MaxSize  int64
}

type Glogger struct {
	LogFile *os.File
	Rander  *rand.Rand
	*Args
}

func OPEN(args *Args) {
	open(args)
}

func CLOSE() {
	close()
}

func DEBUG(format string, v ...interface{}) {
	writeDebug(format, v...)
}

func INFO(format string, v ...interface{}) {
	writeInfo(format, v...)
}

func WARN(format string, v ...interface{}) {
	writeWarn(format, v...)
}

func ERROR(format string, v ...interface{}) {
	writeError(format, v...)
}

func FATAL(format string, v ...interface{}) {
	writeFatal(format, v...)
}

func PANIC(format string, v ...interface{}) {
	writePanic(format, v...)
}

func open(args *Args) {

	close()
	f, err := os.OpenFile(args.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return
	}

	if args.MaxSize <= 0 {
		args.MaxSize = DEFAULT_MAXLOGSIZE
	}

	glogger = &Glogger{
		LogFile: f,
		Rander:  rand.New(rand.NewSource(time.Now().UnixNano())),
		Args:    args,
	}
	SetOutput(f)
	ParseLevel(args.Level)
}

func close() {

	if glogger != nil && glogger.LogFile != nil {
		glogger.LogFile.Close()
		glogger.LogFile = nil
	}
}

func swapLogger() {

	if glogger != nil && glogger.LogFile != nil {

		finfo, err := glogger.LogFile.Stat()
		if err != nil || finfo.Size() < glogger.Args.MaxSize {
			return
		}

		close()
		fpath := filepath.Dir(glogger.Args.FileName) + "\\"
		fname := time.Now().Format("20060102-150405-") + strconv.Itoa(glogger.Rander.Intn(10000)) + ".log"
		os.Rename(glogger.Args.FileName, fpath+fname)
		open(glogger.Args)
	}
}

func writeDebug(format string, v ...interface{}) {
	Debug(format, v...)
	swapLogger()
}

func writeInfo(format string, v ...interface{}) {
	Info(format, v...)
	swapLogger()
}

func writeWarn(format string, v ...interface{}) {
	Warn(format, v...)
	swapLogger()
}

func writeError(format string, v ...interface{}) {
	Error(format, v...)
	swapLogger()
}

func writeFatal(format string, v ...interface{}) {
	Fatal(format, v...)
	swapLogger()
}

func writePanic(format string, v ...interface{}) {
	Panic(format, v...)
	swapLogger()
}
