package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	ft "github.com/x-cray/logrus-prefixed-formatter"

	"notifier/config"
	"notifier/constant"
	"notifier/factory"
)

var Version = "0.0.0"

func main() {
	l := logrus.New()
	l.Level = logrus.DebugLevel
	l.Formatter = &ft.TextFormatter{
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: constant.TimeStampFormat,
	}

	conf, err := config.NewConfig()
	if err != nil {
		l.Fatalf("Error in generating config: %s", err)
	}

	// this is dummy server to listen for http probes on gcp
	go func() {
		fmt.Println(http.ListenAndServe("0.0.0.0:8080", nil))
	}()

	f := factory.NewFactory(l, conf)
	l.Infof("Running notifier server version: %s", Version)
	runner := f.Runner()
	runner.Run()
}
