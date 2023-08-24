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

	lineController := line.NewLineController()
	authController := line_auth.NewLineAuth()

	facades.Route().Prefix("linechat").Group(func(mainRoute route.Route) {
		mainRoute.Prefix("message").Group(func(route route.Route) {
			route.Get("/", userController.Show)
			route.Post("/webhook", lineController.LineWebhookHandler)
			route.Post("/createLineConfig", lineController.CreateLineConfig)
			route.Put("/updateLineConfig", lineController.UpdateLineConfig)
		})

		mainRoute.Prefix("auth").Group(func(routeApiInGroup route.Route) {
			routeApiInGroup.Get("/", authController.LineLogin)
			routeApiInGroup.Get("/callback", authController.LineLoginCallback)
			routeApiInGroup.Get("/refresh", authController.LineRefresh)
			routeApiInGroup.Get("/profile", authController.LineLoginProfile)
			routeApiInGroup.Get("/revoke", authController.LineRevoke)
			routeApiInGroup.Get("/verify", authController.LineVerify)
		})
	})

	// facades.Route().Get("/auth/line", authController.LineLogin)
}
