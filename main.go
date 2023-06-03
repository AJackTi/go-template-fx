package main

import (
	"github.com/AJackTi/go-template-fx/bundlefx"
	"github.com/AJackTi/go-template-fx/httphandler"
	"go.uber.org/fx"
)

// ApplicationConfig ...
type ApplicationConfig struct {
	Address string `yaml:"address"`
}

// Config ...
type Config struct {
	ApplicationConfig `yaml:"application"`
}

func main() {
	fx.New(
		bundlefx.Module,
		fx.Invoke(httphandler.New),
	).Run()
}
