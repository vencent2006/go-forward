package main

import "github.com/sirupsen/logrus"

func Demo() {
	logrus.Info("I am demo")
}
func main() {

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// do demo
	Demo()
}
