package logging

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger interface {
	Info(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
}

type ContentLog struct {
	Buffer []byte
}

type logger struct {
	appName   string
	file      *os.File
	pool      sync.Pool
	logChan   chan *ContentLog
	closeChan chan struct{}
	wg        sync.WaitGroup
}

func (l *logger) Close() error {
	close(l.closeChan)
	l.wg.Wait()
	close(l.logChan)
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

func (l *logger) write(buffer []byte) {
	if _, err := l.file.Write(buffer); err != nil {
		if errClose := l.Close(); errClose != nil {
			fmt.Println(errClose)
		}
		fmt.Println(err)
		return
	}
	os.Stdout.Write(buffer)
}

func (l *logger) processLog() {
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

func (l *logger) cleanBuffer(contentLog *ContentLog) {
	contentLog.Buffer = contentLog.Buffer[:0]
}

func (l *logger) applyLogFormat(format *string, level string) {
	var currentTime string = time.Now().Format("2006-01-02 15:04:05")
	*format = fmt.Sprintf("[%s] %s %s -> %s\n", l.appName, level, currentTime, *format)
}

func (l *logger) Info(format string, v ...any) {
	var contentLog *ContentLog = l.pool.Get().(*ContentLog)
	l.cleanBuffer(contentLog)
	l.applyLogFormat(&format, "Info")
	contentLog.Buffer = fmt.Appendf(contentLog.Buffer, format, v...)
	l.logChan <- contentLog
}

func (l *logger) Warn(format string, v ...any) {
	var contentLog *ContentLog = l.pool.Get().(*ContentLog)
	l.cleanBuffer(contentLog)
	l.applyLogFormat(&format, "Warn")
	contentLog.Buffer = fmt.Appendf(contentLog.Buffer, format, v...)
	l.logChan <- contentLog
}

func (l *logger) Error(format string, v ...any) {
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
) (*logger, error) {
	var flag int = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(path, flag, 0666)
	if err != nil {
		return nil, err
	}
	var loggerPtr *logger = &logger{
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
