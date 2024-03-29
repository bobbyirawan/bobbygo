package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	Logger interface {
		Info(...interface{})
		Infof(string, ...interface{})
		Debug(...interface{})
		Debugf(string, ...interface{})
		Error(...interface{})
		Errorf(string, ...interface{})
		Warning(...interface{})
		Warningf(string, ...interface{})
		Fatal(...interface{})
		Fatalf(string, ...interface{})
		Print(...interface{})
		Printf(string, ...interface{})
		Println(...interface{})
		Instance() interface{}
	}

	Level     string
	Formatter string

	Option struct {
		Level                       Level
		LogFilePath                 string
		Formatter                   Formatter
		MaxSize, MaxBackups, MaxAge int
		Compress                    bool
	}

	lumberjackHook struct {
		lbj    *lumberjack.Logger
		logrus *logrus.Logger
	}

	impl struct {
		instance *logrus.Logger
	}
)

const (
	Info  Level = "INFO"
	Debug Level = "DEBUG"
	Error Level = "ERROR"

	JSONFormatter Formatter = "JSON"
	TextFormatter Formatter = "TEXT"
)

func (l *impl) Info(args ...interface{}) {
	l.instance.Info(args...)
}

func (l *impl) Infof(format string, args ...interface{}) {
	l.instance.Infof(format, args...)
}

func (l *impl) Debug(args ...interface{}) {
	l.instance.Debug(args...)
}

func (l *impl) Debugf(format string, args ...interface{}) {
	l.instance.Debugf(format, args...)
}

func (l *impl) Error(args ...interface{}) {
	l.instance.Error(args...)
}

func (l *impl) Errorf(format string, args ...interface{}) {
	l.instance.Errorf(format, args...)
}

func (l *impl) Warning(args ...interface{}) {
	l.instance.Warning(args...)
}

func (l *impl) Warningf(format string, args ...interface{}) {
	l.instance.Warningf(format, args...)
}

func (l *impl) Fatal(args ...interface{}) {
	l.instance.Fatal(args...)
}

func (l *impl) Fatalf(format string, args ...interface{}) {
	l.instance.Fatalf(format, args...)
}

func (l *impl) Print(args ...interface{}) {
	l.instance.Print(args...)
}

func (l *impl) Println(args ...interface{}) {
	l.instance.Println(args...)
}

func (l *impl) Printf(format string, args ...interface{}) {
	l.instance.Printf(format, args...)
}

func (l *impl) Instance() interface{} {
	return l.instance
}
func New(option *Option) (Logger, error) {
	instance := logrus.New()

	if option.Level == Info {
		instance.Level = logrus.InfoLevel
	}

	if option.Level == Debug {
		instance.Level = logrus.DebugLevel
	}

	if option.Level == Error {
		instance.Level = logrus.ErrorLevel
	}

	var formatter logrus.Formatter

	if option.Formatter == JSONFormatter {
		formatter = &logrus.JSONFormatter{}
	} else {
		formatter = &logrus.TextFormatter{}
	}

	instance.Formatter = formatter

	// - check if log file path does exists
	if option.LogFilePath != "" {
		lbj := &lumberjack.Logger{
			Filename:   option.LogFilePath,
			MaxSize:    option.MaxSize,
			MaxAge:     option.MaxAge,
			MaxBackups: option.MaxBackups,
			LocalTime:  true,
			Compress:   option.Compress,
		}

		instance.Hooks.Add(&lumberjackHook{
			lbj:    lbj,
			logrus: instance,
		})
	}

	return &impl{instance}, nil
}

func (l *lumberjackHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.InfoLevel, logrus.DebugLevel, logrus.ErrorLevel}
}

func (l *lumberjackHook) Fire(entry *logrus.Entry) error {
	_, err := l.logrus.Formatter.Format(entry)

	if err != nil {
		return err
	}

	// if _, err := l.lbj.Write(b); err != nil {
	// 	return err
	// }

	return nil
}
