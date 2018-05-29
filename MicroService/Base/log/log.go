package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/mfslog/goPractice/MicroService/Base/common"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"os"
)

var Logger *zap.Logger


func SettupingLogger(){
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	//topicDebugging := zapcore.AddSync(ioutil.Discard)
	//topicErrors := zapcore.AddSync(ioutil.Discard)

	//kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	//create errorFile
	var slash string = string(os.PathSeparator)
	var logDir string = common.ApplicationDir +slash+"log" + slash
	_,err := os.Open(logDir)

	if err != nil{
		os.MkdirAll(logDir,os.ModePerm)
	}

	errorFile,cls,err := zap.Open(logDir + common.ApplicationName+"_error.log")
	if err != nil{
		os.Exit(1)
	}

	//create debugging
	debugFile,cls, err := zap.Open( logDir+ common.ApplicationName + "_debug.log")
	if err != nil{
		cls()
		os.Exit(1)
	}
	core := zapcore.NewTee(
		//zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder,errorFile , highPriority),
		//zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder,debugFile, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	Logger = zap.New(core)
	defer Logger.Sync()
	grpc_zap.ReplaceGrpcLogger(Logger)
	Logger.Info(common.GetVersionInfo())
	Logger.Error(common.GetVersionInfo())
}