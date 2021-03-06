package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

//LogLevel is log level.
type LogLevel int

//LogLevel enums.
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

const (
	color_red     = uint8(iota + 91) //\x1b[91m%s\x1b[0m
	color_green                      //\x1b[92m%s\x1b[0m
	color_yellow                     //\x1b[93m%s\x1b[0m
	color_blue                       //\x1b[94m%s\x1b[0m
	color_magenta                    //洋红  //\x1b[95m%s\x1b[0m
)

//Logger is a common console and file logger.
type Logger struct {
	Type  string //
	Src   string
	Level LogLevel //Log level
}

//NewLogger will return Logger strcut instance.
func NewLogger(typpe, src string, lv LogLevel) *Logger {
	return &Logger{Type: typpe, Src: src, Level: lv}
}

//LogFatal will record fatal.
func (lgr *Logger) LogFatal(tag, msg interface{}) {
	if lgr.Level <= FATAL {
		lgr.logEvyThg(fmt.Sprintf("\x1b[91m%s\x1b[0m", "FATAL"), tag, msg)
	}
}

//LogError will record error.
func (lgr *Logger) LogError(tag, msg interface{}) {
	if lgr.Level <= ERROR {
		lgr.logEvyThg(fmt.Sprintf("\x1b[95m%s\x1b[0m", "ERROR"), tag, msg)
	}
}

//LogWarn will record warning.
func (lgr *Logger) LogWarn(tag, msg interface{}) {
	if lgr.Level <= WARN {
		lgr.logEvyThg(fmt.Sprintf("\x1b[93m%s\x1b[0m", "WARN"), tag, msg)
	}
}

//LogInfo will record infomation.
func (lgr *Logger) LogInfo(tag, msg interface{}) {
	if lgr.Level <= INFO {
		lgr.logEvyThg(fmt.Sprintf("\x1b[92m%s\x1b[0m", "INFO"), tag, msg)
	}
}

//LogDebug will record debug infomation.
func (lgr *Logger) LogDebug(tag, msg interface{}) {
	if lgr.Level <= DEBUG {
		lgr.logEvyThg("DEBUG", tag, msg)
	}
}

func (lgr *Logger) logEvyThg(lv string, tag, msg interface{}) {
	spath := fmt.Sprintf("./Log/%s/%s", lgr.Type, time.Now().Format("2006-01-02"))
	os.MkdirAll(spath, os.ModePerm)
	f, err := os.OpenFile(fmt.Sprintf("%s/%s.txt", spath, lgr.Src), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0665)
	if err != nil {
		log.Println("failed to open log file: " + err.Error())
	}
	defer f.Close()
	log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Printf("[%s] %v: %v%s", lv, tag, msg, NewLine())
}

func NewLine() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else if runtime.GOOS == "darwin" {
		return "\r"
	}
	return "\n"
}
