package cmd

import (
	"testing/configs"
	"testing/routes"
)

// biasanya digunakan untuk inject dependencies injection

func Run() {
	configs.InitDB(configs.SetENV())
	configs.InitCache()

	routes.Routes()
}
