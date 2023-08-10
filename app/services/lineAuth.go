package services

import (
	"goravel/app/modules/line_auth"

	"github.com/goravel/framework/facades"
)

var (
	LineAuth *line_auth.LineLogin
)

func LineAuthSetup() {
	config := facades.Config()
	config.Add("line_oauth", map[string]any{
		"LINE_CLIENT_ID":     config.Env("LINE_CLIENT_ID", ""),
		"LINE_CLIENT_SECRET": config.Env("LINE_CLIENT_SECRET", ""),
		"LINE_REDIRECT_URL":  config.Env("LINE_REDIRECT_URL", ""),
	})
	var (
		CLIENT_ID     = config.GetString("line_oauth.LINE_CLIENT_ID")
		CLIENT_SECRET = config.GetString("line_oauth.LINE_CLIENT_SECRET")
		REDIRECT_URL  = config.GetString("line_oauth.LINE_REDIRECT_URL")
	)
	lineAuth := line_auth.NewLineLogin(CLIENT_ID, CLIENT_SECRET, REDIRECT_URL)
	LineAuth = lineAuth
}
