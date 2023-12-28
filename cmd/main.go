package main

import (
	"fmt"
	"go-crud-v2/internal/config"
)

func main() {
	viper := config.NewViper()
	db := config.NewGorm(viper)
	log := config.NewLogger(viper)
	validate := config.NewValidator(viper)
	app := config.NewFiber(viper)

	conf := config.BootstrapConfig{
		DB:       db,
		Log:      log,
		App:      app,
		Validate: validate,
		Viper:    viper,
	}
	config.Bootstrap(&conf)
	webPort := viper.GetInt("web.port")

	if err := app.Listen(fmt.Sprintf(":%d", webPort)); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
