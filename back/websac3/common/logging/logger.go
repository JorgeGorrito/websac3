package logging

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type ContentLog struct {
	Buffer []byte
}

type Logger struct {
	appName   string
	file      *os.File
	pool      sync.Pool
	logChan   chan *ContentLog
	closeChan chan struct{}
	wg        sync.WaitGroup
}

func (l *Logger) Close() error {
	close(l.closeChan)
	l.wg.Wait()
	close(l.logChan)
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

func (l *Logger) write(buffer []byte) {
	if _, err := l.file.Write(buffer); err != nil {
		if errClose := l.Close(); errClose != nil {
			fmt.Println(errClose)
		}
		fmt.Println(err)
		return
	}
	os.Stdout.Write(buffer)
}

func (l *Logger) processLog() {
	defer l.wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Logger panic: %v\n", r)
		}
	}()
	for {
		select {
		case <-l.closeChan:
			return
		case log := <-l.logChan:
			l.write(log.Buffer)
			l.pool.Put(log)
		}
	}
}

func (l *Logger) cleanBuffer(contentLog *ContentLog) {
	contentLog.Buffer = contentLog.Buffer[:0]
}

func (l *Logger) applyLogFormat(format *string, level string) {
	var currentTime string = time.Now().Format("2006-01-02 15:04:05")
	*format = fmt.Sprintf("[%s] %s %s -> %s\n", l.appName, level, currentTime, *format)
}

func (l *Logger) Info(format string, v ...any) {
	var contentLog *ContentLog = l.pool.Get().(*ContentLog)
	l.cleanBuffer(contentLog)
	l.applyLogFormat(&format, "Info")
	contentLog.Buffer = fmt.Appendf(contentLog.Buffer, format, v...)
	l.logChan <- contentLog
}

func (l *Logger) Warn(format string, v ...any) {
	var contentLog *ContentLog = l.pool.Get().(*ContentLog)
	l.cleanBuffer(contentLog)
	l.applyLogFormat(&format, "Warn")
	contentLog.Buffer = fmt.Appendf(contentLog.Buffer, format, v...)
	l.logChan <- contentLog
}

func (l *Logger) Error(format string, v ...any) {
	var contentLog *ContentLog = l.pool.Get().(*ContentLog)
	l.cleanBuffer(contentLog)
	l.applyLogFormat(&format, "Error")
	contentLog.Buffer = fmt.Appendf(contentLog.Buffer, format, v...)
	l.logChan <- contentLog
}

func NewLogger(
	appName string,
	path string,
	bufferSize int,
	maxJobs int,
) (*Logger, error) {
	var flag int = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(path, flag, 0666)
	if err != nil {
		return nil, err
	}
	var loggerPtr *Logger = &Logger{
		appName:   appName,
		file:      file,
		closeChan: make(chan struct{}),
		logChan:   make(chan *ContentLog, maxJobs),
		pool: sync.Pool{
			New: func() any {
				return &ContentLog{
					Buffer: make([]byte, 0, bufferSize),
				}
			},
		},
	}
	loggerPtr.wg.Add(1)
	go loggerPtr.processLog()
	return loggerPtr, nil
}
