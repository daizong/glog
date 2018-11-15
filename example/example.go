package main

import (
	"log"
	"os"

	"github.com/daizong/glog"
)

func main() {
	logger := glog.NewLogger(os.Stdout, "", log.LstdFlags|log.Lshortfile, glog.DEBUG)
	logger.SetCallDepth(2)
	logger.SetLevel(glog.INFO)
	logger.Debugf("current level: %s, %s", logger.GetLevelString(), "this a debug log")
	logger.Infof("current level: %s, %s", logger.GetLevelString(), "this a info log")
	logger.Warningf("current level: %s, %s", logger.GetLevelString(), "this a warning log")
	logger.Errorf("current level: %s, %s", logger.GetLevelString(), "this a error log")
	logger.Fatalf("current level: %s, %s", logger.GetLevelString(), "this a fatal log")
}
