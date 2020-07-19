package strategy

import "testing"

func TestNewLogManager(t *testing.T) {
	filelog := &FileLoging{}
	logManager := NewLogManager(filelog)
	logManager.Info()
	logManager.Error()


	//
	dblog := &DataBaseLogging{}
	logManager.Logging = dblog
	logManager.Info()
	logManager.Error()


}

func TestFileLoging_Info(t *testing.T) {

}