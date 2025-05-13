package pkg

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.StandardLogger()
	logger.AddHook(&StackHook{})
	// logger.SetFormatter(&logrus.TextFormatter{})

}

type StackHook struct{}

func (s *StackHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}
func (s *StackHook) Fire(entry *logrus.Entry) error {

	if e, ok := entry.Data[logrus.ErrorKey]; ok {
		if v, ok := e.(StackTraceInterface); ok {
			entry.Data["Stack"] = v.StackTrace()
		}

		if v, ok := e.(error); ok {
			entry.Message = v.Error()
		}

		delete(entry.Data, logrus.ErrorKey)
	}

	return nil
}

func PrintErr(err error) {
	logger.WithError(err).Error()
}
