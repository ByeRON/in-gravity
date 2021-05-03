package cmd

import (
	"flag"
	"in-gravity/config"
	"in-gravity/ui/restapi"
)

var ServeCmd = flag.NewFlagSet("serve", flag.ExitOnError)
var confPath = ServeCmd.String("config", "", "please set application config path")

func Serve(args []string) {
	err := ServeCmd.Parse(args)
	if err != nil {
		panic(err)
	}
	appConf, err := config.ApplicationConfigFromYAML(confPath)
	if err != nil {
		panic(err)
	}
	restapi.Serve(appConf)
}
