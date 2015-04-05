package transform

import (
	"da"
	"log"
	"path/filepath"

	"github.com/hjin-me/banana"
)

type Conf struct {
	MySQL string `yaml:"mysql"`
}

func install() {
	f, err := filepath.Abs("conf")
	if err != nil {
		panic(err)
	}
	banana.SetBaseDir(f)

	cfg := Conf{}

	_, err = banana.Config("app.yaml", &cfg)
	if err != nil {
		panic(err)
	}
	err = da.Create(cfg.MySQL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("database connected")
	err = importUsers()
	if err != nil {
		panic(err)
	}
	defer da.Close()
}
