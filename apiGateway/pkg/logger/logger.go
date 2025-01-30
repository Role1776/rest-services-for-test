package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logger

type logger struct {
	*logrus.Entry
}

func init() {
	logrusNew := logrus.New()
	logrusNew.SetReportCaller(true)

	logrusNew.Formatter = &logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableColors:          false,
		ForceQuote:             true,
		DisableLevelTruncation: true,
	}
	log = &logger{logrus.NewEntry(logrusNew)}
}

func Log() *logger {
	return log
}

func (l *logger) LogWithField(k string, v interface{}) *logger {
	return &logger{l.WithField(k, v)}
}
