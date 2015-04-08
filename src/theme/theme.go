package theme

import (
	"fmt"
	"sync"

	"github.com/hjin-me/banana"
)

type Config struct {
	CP string `yaml:"control_panel"`
	UI string `yaml:"user_interface"`
}

var globalCfg = Config{}
var once sync.Once

func loadCfg() Config {
	once.Do(func() {
		banana.Config("theme.yaml", &globalCfg)
	})
	return globalCfg
}

func UI(page string) string {
	cfg := loadCfg()
	return fmt.Sprintf("%s:page/%s", cfg.UI, page)
}

func CP(page string) string {
	cfg := loadCfg()
	return fmt.Sprintf("%s:page/%s", cfg.CP, page)
}
