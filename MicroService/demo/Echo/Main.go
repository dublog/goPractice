package Echo

import (
	"github.com/spf13/cobra"
	"fmt"
)

func runServer(){
	fmt.Println("Run Server")
}

func Main(version string){

	var configFile = "./config/config.json"
	var rootCmd = &cobra.Command{
		Use: "",
		Short: "Echo",
		Long: "Echo Server",
		Run:func(cmd *cobra.Command, args []string){
			runServer()
		},
	}

	rootCmd.Flags().StringVarP(&configFile,"config", "c","config.json","path to config file")

	var versionCmd = &cobra.Command{
		Use:"version",
		Short: "v",
		Long:"show version",
		Run:func(cmd *cobra.Command, args []string){
			fmt.Print(version)
		},
	}
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}
