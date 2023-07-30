package main

import (
	"fmt"
	"github.com/kozhamseitova/api-blog/internal/app"
	"github.com/kozhamseitova/api-blog/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#v", cfg))

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
