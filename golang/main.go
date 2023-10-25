package main

import (
	"bitbucket.org/viasoftkorp/korp.sdk/auto_update"
	"bitbucket.org/viasoftkorp/korp.sdk/consul"
	"bitbucket.org/viasoftkorp/korp.sdk/fiber_helper"
	"bitbucket.org/viasoftkorp/korp.sdk/service_discovery"
	"bitbucket.org/viasoftkorp/korp.sdk/service_info"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	service_info.ServiceName = "Samples"

	consul.AddKorpConsul(consul.KorpConsulConfig{})
	auto_update.MustDoUpdate()

	app := fiber_helper.NewFiberApp(fiber.Config{
		BodyLimit: int(10.0) * 1024 * 1024,
	})
	app.Use(logger.New())
	service_discovery.AddKorpFiberHealthCheck(app)

	initRoutes(app)

	service_discovery.MustAddServiceDiscovery(service_discovery.NewServiceDiscoveryConfig(fiber_helper.NewFiberShutdownServer(app), "puc/cursos/"))

	fiber_helper.Listen(app)
}
