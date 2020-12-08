package main

import (
	"fmt"
	"github.com/hinha/api-box/log"
	"github.com/hinha/api-box/provider/api"
	"github.com/hinha/api-box/provider/auth"
	"github.com/hinha/api-box/provider/command"
	"github.com/hinha/api-box/provider/infrastructure"
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

	// Infra
	infra, err := infrastructure.Fabricate()
	if err != nil {
		panic(err)
	}
	defer infra.Close()

	// Infra DB Mysql
	db, err := infra.DB()
	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	// API
	apiEngine := api.Fabricate(PORT)
	apiEngine.FabricateCommand(cmd)

	o2auth := auth.FabricateGoogle()
	o2auth.FabricateAPI(apiEngine)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
