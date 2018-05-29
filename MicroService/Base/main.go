package main

import (
	"github.com/spf13/cobra"
	"github.com/FZambia/viper-lite"
	"path/filepath"
	"os"
	"fmt"
	"github.com/mfslog/goPractice/MicroService/Base/common"
	"github.com/mfslog/goPractice/MicroService/Base/log"
	"go.uber.org/zap"
)
var(
	configPath string = "config.json"
	logPath string = "./log"
	logLevel string = "info"
	logger *zap.Logger
)
func main(){
	common.ApplicationName = filepath.Base(os.Args[0])
	common.ApplicationDir = filepath.Dir(os.Args[0])
	var rootCmd = &cobra.Command{
		Use:"",
		Short: common.ApplicationName + " command --flag",
		Long: "execute appliation, startup  service",
		Run: func(cmd *cobra.Command, args []string) {

			bindFlags := []string{
				"log_level","config",
			}

			for _,flag := range bindFlags {
				viper.BindPFlag(flag,cmd.Flags().Lookup(flag))
			}
			log.SettupingLogger()
			logger = log.Logger
			logger.Info("load config file [" +viper.GetString("config") + "]")
		},
	}

	var versionCmd = &cobra.Command{
		Use:"version",
		Short:"version information",
		Long:"show version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(common.GetVersionInfo())
		},
	}

	rootCmd.Flags().StringVarP(&configPath,"config","c","config.json","configer file path")
	rootCmd.Flags().StringVarP(&logLevel, "log_level","","info","set the log level:")


	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}