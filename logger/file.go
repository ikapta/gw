package logger

import (
	"log"
	"os"
)

type FileLogger struct {
  FilePath string
  FileName string
  Info *log.Logger
  Warn *log.Logger
  Error *log.Logger
  Write func(msg string)
}

// ConfigDefault is the default config
var ConfigDefault = FileLogger{
  FilePath: "./log/",
  FileName: "gw.log",
}

func Setup(conf ...FileLogger) (FileLogger, error) {

  var cfg = conf[0]

  if (cfg.FilePath == "") {
    cfg.FilePath = ConfigDefault.FilePath
  }

  if (cfg.FileName == "") {
    cfg.FileName = ConfigDefault.FileName
  }

  var fullFilePath = cfg.FilePath + cfg.FileName

  if _, err := os.Stat(fullFilePath); os.IsNotExist(err) {
    os.MkdirAll(cfg.FilePath, os.ModePerm)
  }

  f, err := os.OpenFile(fullFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

  if err != nil {
    log.Printf("create file %s failed: %v", cfg.FilePath, err)
  }

  cfg.Info = log.New(f, "[info] ", log.LstdFlags|log.Lmicroseconds)

  cfg.Warn = log.New(f, "[warn] ", log.LstdFlags|log.Lmicroseconds)

  cfg.Error = log.New(f, "[Error] ", log.LstdFlags|log.Lmicroseconds)

  cfg.Write = func(msg string) {
    cfg.Info.Print(msg)
  }

  return cfg, err
}
