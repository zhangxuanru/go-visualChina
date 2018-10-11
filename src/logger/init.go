package logger

import (
	"time"
	"fmt"
	"os"
	"log"
	"io"
)

const LOG_DIR = "log/"

var(
	Info *log.Logger
	Error *log.Logger
)

func init()  {
	year :=  time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	infoLog := fmt.Sprintf(LOG_DIR+"info-%d-%d-%d.log",year,month,day)
	errorLog := fmt.Sprintf(LOG_DIR+"error-%d-%d-%d.log",year,month,day)

	infoFile, _ := os.OpenFile(infoLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	errorFile, _ := os.OpenFile(errorLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)

	Info = log.New(infoFile,"Info: ",log.LstdFlags | log.Lshortfile | log.LUTC)
	Error = log.New(io.MultiWriter(os.Stderr,errorFile),"error: ",log.LstdFlags | log.Lshortfile | log.LUTC)
}

