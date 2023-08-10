package services

import (
	"github.com/goravel/framework/contracts/http"
)

type ResMessageService struct {
}

type ResMessageImpl struct {
}

func NewResMessageService() *ResMessageImpl {
	return &ResMessageImpl{}
}

func (r *ResMessageImpl) Json(resCode int, Result any) (int, map[string]interface{}) {
	message := http.StatusText(resCode)

	resMessage := map[string]interface{}{
		"status":  resCode,
		"message": message,
		"result":  Result,
	}
	return resCode, resMessage
}
