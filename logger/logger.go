package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func Test() {
	var log = logrus.New()
	log.Out = os.Stdout

	log.Infof("test: %d, %s", 1, "hello world")
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}

func TestLogToFile() {
	log := logrus.New()
	log.SetReportCaller(true)       // 设置记录调用者信息
	log.SetOutput(&logFileWriter{}) // 设置文件输出
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.Infof("test: %d, %s", 1, "hello world")
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}

// 继承io.Writer: 实现Write(p []byte) (n int, err error)
type logFileWriter struct {
	file  *os.File
	ctime time.Time // file创建时间
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if time.Now().Format("2006-01-02") != p.ctime.Format("2006-01-02") {
		p, err = New4Day()

		if err != nil {
			return
		}
	}

	// if p.file == nil {
	// 	return 0, errors.New("logFileWriter:Write file not opened")
	// }

	n, err = p.file.Write(data)
	return
}

// 每天一个文件
func New4Day() (*logFileWriter, error) {
	// ensure dir exists
	os.MkdirAll("tmp/logs", os.ModePerm)
	//
	now := time.Now()
	filename := fmt.Sprintf("./tmp/logs/log_%s.log", now.Format("2006-01-02"))
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		log.Panic("打开日志文件失败:" + err.Error())
		return nil, err
	}

	return &logFileWriter{
		file:  file,
		ctime: now,
	}, nil
}
