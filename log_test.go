package log

import "testing"

func TestLog(t *testing.T) {
	SetLevel(DebugLevel)
	Debugf("debug")
	Infof("info")
	Warnf("warn")
	Errorf("error")
	Fatalf("fatal")
}
