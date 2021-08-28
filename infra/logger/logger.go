package logger

import "github.com/sirupsen/logrus"

type Log struct {
	log *logrus.Logger
}

func NewLogger() Log {
	return Log{log: logrus.New()}
}

func (l Log) Infof(msg string, params ...interface{}) {
	l.log.Infof(msg, params...)
}

func (l Log) Debugf(msg string, params ...interface{}) {
	l.log.Debugf(msg, params...)
}

func (l Log) Warnf(msg string, params ...interface{}) {
	l.log.Warnf(msg, params...)
}
