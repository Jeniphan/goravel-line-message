package routes

import (
	// "fmt"

	// "encoding/json"

	// "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	// "github.com/gorilla/mux"
	"goravel/app/http/controllers"
	"goravel/app/http/controllers/line"
	"goravel/app/http/controllers/line_auth"
	// "log"
	// "net/http"
)

func Web() {

	userController := controllers.NewUserController()
	facades.Route().Get("/", userController.Show)

	lineController := line.NewLineController()
	facades.Route().Prefix("message").Group(func(route route.Route) {
		route.Post("/webhook", lineController.LineWebhookHandler)
		route.Post("/createLineConfig", lineController.CreateLineConfig)
		route.Put("/updateLineConfig", lineController.UpdateLineConfig)
	})

	authController := line_auth.NewLineAuth()
	// facades.Route().Get("/auth/line", authController.LineLogin)
	facades.Route().Prefix("auth/line").Group(func(routeApiInGroup route.Route) {
		routeApiInGroup.Get("/", authController.LineLogin)
		routeApiInGroup.Get("/callback", authController.LineLoginCallback)
		routeApiInGroup.Get("/refresh", authController.LineRefresh)
		routeApiInGroup.Get("/profile", authController.LineLoginProfile)
		routeApiInGroup.Get("/revoke", authController.LineRevoke)
		routeApiInGroup.Get("/verify", authController.LineVerify)
	})
}
