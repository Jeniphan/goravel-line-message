package services

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type LineConfig struct {
}

func NewLineConfigs() *LineConfig {
	return &LineConfig{}
}

func (r *LineConfig) CreateLineConfig(dataLineConfigs models.LineConfigs) (any, error) {
	// Check LineChannel Id form Database
	err := facades.Orm().Query().FindOrFail(&dataLineConfigs, "line_channel_id=?", dataLineConfigs.LineChannelId)
	if err != nil {
		// return nil, err
		err = facades.Orm().Query().Create(&dataLineConfigs)

		if err != nil {
			return nil, err
		} else {
			return dataLineConfigs, nil
		}
	}

	return "Line Channel Id is taken", nil
}

func (r *LineConfig) UpdateLineConfig(lineChannelId string, dataUpdate models.LineConfigs) (any, error) {
	result, err := facades.Orm().Query().Model(models.LineConfigs{}).Where("line_channel_id", lineChannelId).Update(dataUpdate)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
