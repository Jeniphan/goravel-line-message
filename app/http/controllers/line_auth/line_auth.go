package line_auth

import (
	"fmt"
	"goravel/app/models"
	"goravel/app/services"
	"time"

	"github.com/goravel/framework/contracts/http"

	"github.com/goravel/framework/facades"
)

type LineAuth struct {
	//Dependent services
}

func NewLineAuth() *LineAuth {
	return &LineAuth{
		//Inject services
	}
}

func (r *LineAuth) Index(ctx http.Context) {
}

func (r *LineAuth) LineLogin(ctx http.Context) {
	validator, err := ctx.Request().Validate(map[string]string{
		"redirect_url": "required",
	})
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	if validator.Fails() {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, validator.Errors().All()))
		return
	}
	var input struct {
		RedirectUri string `json:"redirect_url"`
	}
	if err := validator.Bind(&input); err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	lineAuth := services.LineAuth
	ctx.Response().Redirect(http.StatusTemporaryRedirect, lineAuth.GetLink(input.RedirectUri))
}

func (r *LineAuth) LineLoginCallback(ctx http.Context) {
	validator, err := ctx.Request().Validate(map[string]string{
		"code":  "required",
		"state": "required",
		// "redirect_url": "required",
	})
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	if validator.Fails() {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, validator.Errors().All()))
		return
	}
	var input struct {
		Code  string `json:"code"`
		State string `json:"state"`
		// RedirectUrl string `json:"redirect_url"`
	}
	if err := validator.Bind(&input); err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	lineAuth := services.LineAuth
	token := lineAuth.Token(input.Code, input.State)
	// fmt.Println(token)
	if token.AccessToken != "" {
		if err := facades.Orm().Query().Create(&models.Logs{
			IP_Address:   ctx.Request().Ip(),
			Content:      fmt.Sprintf("access_token=%v&expires_in=%v&refresh_token=%v", token.AccessToken, token.ExpiresIn, token.RefreshToken),
			Url_Callback: fmt.Sprintf("%v?access_token=%v&expires_in=%v&refresh_token=%v", input.State, token.AccessToken, token.ExpiresIn, token.RefreshToken),
			LoginAt:      time.Now(),
		}); err != nil {
			fmt.Println(err.Error())
		}
		ctx.Response().Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%v?access_token=%v&expires_in=%v&refresh_token=%v", input.State, token.AccessToken, token.ExpiresIn, token.RefreshToken))
		return
	} else {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, ""))
		return
	}
}

func (r *LineAuth) LineLoginProfile(ctx http.Context) {
	validator, err := ctx.Request().Validate(map[string]string{
		"access_token": "required",
	})
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	if validator.Fails() {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, validator.Errors().All()))
		return
	}
	var input struct {
		Token string `json:"access_token"`
	}
	if err := validator.Bind(&input); err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	lineAuth := services.LineAuth
	profile := lineAuth.Profile(input.Token)
	if profile.UserId != "" {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusOK, profile))
		return
	}
	ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, ""))
}

func (r *LineAuth) LineRefresh(ctx http.Context) {
	validator, err := ctx.Request().Validate(map[string]string{
		"refresh_token": "required",
	})
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	if validator.Fails() {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, validator.Errors().All()))
		return
	}
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := validator.Bind(&input); err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	lineAuth := services.LineAuth
	// fmt.Println(input)
	token := lineAuth.Refresh(input.RefreshToken)
	if token.AccessToken != "" {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusOK, token))
		return
	}
	ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, ""))
}

func (r *LineAuth) LineRevoke(ctx http.Context) {
	validator, err := ctx.Request().Validate(map[string]string{
		"access_token": "required",
	})
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	if validator.Fails() {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, validator.Errors().All()))
		return
	}
	var input struct {
		Token string `json:"access_token"`
	}
	if err := validator.Bind(&input); err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	lineAuth := services.LineAuth
	response := lineAuth.Revoke(input.Token)
	ctx.Response().Json(services.NewResMessageService().Json(http.StatusOK, response))
}

func (r *LineAuth) LineVerify(ctx http.Context) {
	validator, err := ctx.Request().Validate(map[string]string{
		"access_token": "required",
	})
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	if validator.Fails() {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, validator.Errors().All()))
		return
	}
	var input struct {
		Token string `json:"access_token"`
	}
	if err := validator.Bind(&input); err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}
	lineAuth := services.LineAuth
	response := lineAuth.Verify(input.Token)
	ctx.Response().Json(services.NewResMessageService().Json(http.StatusOK, response))
}

func (r *LineAuth) GoogleLogin(ctx http.Context) {

}
func (r *LineAuth) GoogleLoginCallback(ctx http.Context) {

}

func (r *LineAuth) AzureLogin(ctx http.Context) {

}

func (r *LineAuth) AzureLoginCallback(ctx http.Context) {

}
