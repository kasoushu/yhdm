package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type MyFormatter struct{}

func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("[CAI-%s] %s  %s\n",strings.ToUpper(entry.Level.String()),timestamp, entry.Message)
	return []byte(msg), nil
}
func InitLogger() *logrus.Logger  {
	var log = logrus.New()
	log.SetFormatter(new(MyFormatter))
	return log
}