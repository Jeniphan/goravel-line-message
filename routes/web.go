package routes

import (
	// "fmt"

	// "encoding/json"

	// "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	// "github.com/gorilla/mux"
	"goravel/app/http/controllers"
	"goravel/app/http/controllers/line"
	// "log"
	// "net/http"
)

func Web() {

	userController := controllers.NewUserController()
	facades.Route().Get("/", userController.Show)

	lineController := line.NewLineController()
	facades.Route().Post("/webhook", lineController.LineWebhookHandler)

	facades.Route().Get("/testDB", lineController.TestDb)
}
