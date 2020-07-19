package strategy

import "fmt"

//日志管理器
type LogManager struct {
	Logging
}
func NewLogManager(loging Logging) *LogManager{
	return &LogManager{loging}
}

//日志接口
type Logging interface {
	Info()
	Error()
}

//
type FileLoging struct {}

func (fl *FileLoging) Info()  {
	fmt.Println("file Info")
}
func (fl *FileLoging)Error()  {
	fmt.Println("file Error")
}

type DataBaseLogging struct {}

func (dl *DataBaseLogging)Info()  {
	fmt.Println("database Info")
}
func (dl *DataBaseLogging)Error()  {
	fmt.Println("database Error")
}

