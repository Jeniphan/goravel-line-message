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
	facades.Route().Post("api/webhook", lineController.LineWebhookHandler)

	authController := line_auth.NewLineAuth()
	facades.Route().Get("/auth/line", authController.LineLogin)
	facades.Route().Prefix("api").Group(func(routeApiInGroup route.Route) {
		routeApiInGroup.Prefix("auth").Group(func(routeAuth route.Route) {
			routeAuth.Prefix("line").Group(func(routeAuthLine route.Route) {
				routeAuthLine.Get("/callback", authController.LineLoginCallback)
				routeAuthLine.Get("/refresh", authController.LineRefresh)
				routeAuthLine.Get("/profile", authController.LineLoginProfile)
				routeAuthLine.Get("/revoke", authController.LineRevoke)
				routeAuthLine.Get("/verify", authController.LineVerify)
			})
		})
	})

}
