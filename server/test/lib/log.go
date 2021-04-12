package lib

import (
	"github.com/sirupsen/logrus"
	"os"
)

var LLog *logrus.Logger

func init() {
	LLog =logrus.New()
	LLog.SetFormatter(&logrus.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	LLog.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	LLog.SetLevel(logrus.TraceLevel)
	LLog.SetReportCaller(true)
	//LLog.Infof(time.Now().String())
}
