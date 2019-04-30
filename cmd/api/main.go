package main

import (
	"github.com/checkaayush/amigo/api"
	"github.com/checkaayush/amigo/pkg/config"
)

const AppName = "amigo"

func main() {
	cfg, err := config.Load(AppName)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
