package main

import (
	"github.com/hinha/api-box/log"
	"github.com/hinha/api-box/provider/api"
	"github.com/hinha/api-box/provider/command"
	"github.com/hinha/api-box/provider/oauth"
	"github.com/subosito/gotenv"
	"os"
	"time"
)

const PORT = 8000

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
	loc, _ := time.LoadLocation(os.Getenv("TZ"))
	time.Local = loc
	_ = gotenv.Load()
	log.NewLogger()
}

func main() {
	cmd := command.Fabricate()

	//infra := infrastructure.Fabricate()
	//// Infra
	//network := infra.GoogleOAuthNetwork()
	//network.GET()

	// API
	apiEngine := api.Fabricate(PORT)
	apiEngine.FabricateCommand(cmd)

	o2auth := oauth.FabricateGoogle()
	o2auth.FabricateAPI(apiEngine)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
