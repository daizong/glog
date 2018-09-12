package glog

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSetOutput(t *testing.T) {
	logger := NewLogger(os.Stdout, "", log.LstdFlags, DEBUG)
	logger.Debugf("std out put")
	logger.SetOutput(os.Stderr)
	logger.Debugf("std err put")
}

func TestGetSetLevel(t *testing.T) {
	logger := NewLogger(os.Stdout, "", log.LstdFlags, DEBUG)
	if logger.GetLevelString() != DEBUG {
		t.Fail()
	}
	logger.SetLevel(ERROR)
	if logger.GetLevelString() != ERROR {
		t.Fail()
	}
}

func TestGlogUsing(t *testing.T) {
	logger := NewLogger(os.Stdout, "", log.LstdFlags, DEBUG)

	logger.Debugf("current level: %s, %s", logger.GetLevelString(), "this a debug log")
	logger.Infof("current level: %s, %s", logger.GetLevelString(), "this a info log")
	logger.Warningf("current level: %s, %s", logger.GetLevelString(), "this a warning log")
	logger.Errorf("current level: %s, %s", logger.GetLevelString(), "this a error log")
	logger.Fatalf("current level: %s, %s", logger.GetLevelString(), "this a fatal log")

	fmt.Println("================================================================")
	logger.SetLevel(INFO)
	logger.Debugf("current level: %s, %s", logger.GetLevelString(), "this a debug log")
	logger.Infof("current level: %s, %s", logger.GetLevelString(), "this a info log")
	logger.Warningf("current level: %s, %s", logger.GetLevelString(), "this a warning log")
	logger.Errorf("current level: %s, %s", logger.GetLevelString(), "this a error log")
	logger.Fatalf("current level: %s, %s", logger.GetLevelString(), "this a fatal log")

	fmt.Println("================================================================")
	logger.SetLevel(WARNING)
	logger.Debugf("current level: %s, %s", logger.GetLevelString(), "this a debug log")
	logger.Infof("current level: %s, %s", logger.GetLevelString(), "this a info log")
	logger.Warningf("current level: %s, %s", logger.GetLevelString(), "this a warning log")
	logger.Errorf("current level: %s, %s", logger.GetLevelString(), "this a error log")
	logger.Fatalf("current level: %s, %s", logger.GetLevelString(), "this a fatal log")

	fmt.Println("================================================================")
	logger.SetLevel(ERROR)
	logger.Debugf("current level: %s, %s", logger.GetLevelString(), "this a debug log")
	logger.Infof("current level: %s, %s", logger.GetLevelString(), "this a info log")
	logger.Warningf("current level: %s, %s", logger.GetLevelString(), "this a warning log")
	logger.Errorf("current level: %s, %s", logger.GetLevelString(), "this a error log")
	logger.Fatalf("current level: %s, %s", logger.GetLevelString(), "this a fatal log")

	fmt.Println("================================================================")
	logger.SetLevel(FATAL)
	logger.Debugf("current level: %s, %s", logger.GetLevelString(), "this a debug log")
	logger.Infof("current level: %s, %s", logger.GetLevelString(), "this a info log")
	logger.Warningf("current level: %s, %s", logger.GetLevelString(), "this a warning log")
	logger.Errorf("current level: %s, %s", logger.GetLevelString(), "this a error log")
	logger.Fatalf("current level: %s, %s", logger.GetLevelString(), "this a fatal log")

	fmt.Println("================================================================")
}
