package line

import (
	"fmt"
	"goravel/app/models"
	"goravel/app/services"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	myHttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/line/line-bot-sdk-go/linebot"
)

type createLineConfig struct {
	LineChanelId           string `json:"line_channel_id"`
	LineChannelName        string `json:"line_chanel_name"`
	LineChannelSecret      string `json:"line_chanel_secret"`
	LineChannelAccessToken string `json:"line_channel_access_token"`
}

type updateLineConfig struct {
	LineChannelName        string `json:"line_chanel_name"`
	LineChannelSecret      string `json:"line_chanel_secret"`
	LineChannelAccessToken string `json:"line_channel_access_token"`
}

var lineConfigsService = services.NewLineConfigs()

type LineController struct {
	//Dependent services
}

func NewLineController() *LineController {
	return &LineController{
		//Inject services
	}
}

func (r *LineController) LineWebhookHandler(ctx myHttp.Context) {

	lineChannelId := ctx.Request().Query("lineChannelId")

	if lineChannelId == "" {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, "Line channel Id is null"))
		return
	}

	var LineConfig models.LineConfigs

	err := facades.Orm().Query().Find(&LineConfig, "line_channel_id=?", lineChannelId)

	fmt.Println(LineConfig)
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
		return
	}

	// config := facades.Config()
	bot, err := linebot.New(
		LineConfig.LineChannelSecret,
		LineConfig.LineChannelAccessToken,
	)

	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
		return
	}

	// Line := new(LineDistination)
	// err = ctx.Request().Bind(Line)

	events, err := bot.ParseRequest(ctx.Request().Origin())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err))
		} else {
			ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
		}
		return
	}

	for _, event := range events {
		userId, err := handleUserLine(event.Source.UserID, string(event.Source.Type))

		if err != nil {
			ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
			return
		}

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.ImageMessage:
				// Handle image message
				imageId, err := handleImageMessage(bot, message, userId, LineConfig.Id)
				if err != nil {
					ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
					return
				}
				handleRepliedMessage(bot, event.ReplyToken, "image", *imageId, *userId, LineConfig.Id)

			case *linebot.TextMessage:
				textId, err := handleTextMessage(message, userId, LineConfig.Id)
				if err != nil {
					ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
					return
				}
				handleRepliedMessage(bot, event.ReplyToken, "text", *textId, *userId, LineConfig.Id)

			case *linebot.StickerMessage:
				stickerId, err := handleStickerMessage(message, userId, LineConfig.Id)
				if err != nil {
					ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
					return
				}
				handleRepliedMessage(bot, event.ReplyToken, "sticker", *stickerId, *userId, LineConfig.Id)

			case *linebot.LocationMessage:
				locationId, err := handleLocationMessage(message, userId, LineConfig.Id)
				if err != nil {
					ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
					return
				}
				handleRepliedMessage(bot, event.ReplyToken, "location", *locationId, *userId, LineConfig.Id)
			case *linebot.AudioMessage:
				audioId, err := handleAudioMessage(bot, message, userId, LineConfig.Id)
				if err != nil {
					ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
					return
				}
				handleRepliedMessage(bot, event.ReplyToken, "audio", *audioId, *userId, LineConfig.Id)
			default:
				// Handle other message types
				// ...
			}
		}
	}

}

func (r *LineController) CreateLineConfig(ctx myHttp.Context) {

	// var lineConfigsService = services.NewLineConfigs()
	dataCreate := new(createLineConfig)
	err := ctx.Request().Bind(&dataCreate)
	fmt.Println(err)
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err.Error()))
		return
	}

	var lineConfigs models.LineConfigs

	lineConfigs.LineChannelId = dataCreate.LineChanelId
	lineConfigs.LineChannelName = dataCreate.LineChannelName
	lineConfigs.LineChannelSecret = dataCreate.LineChannelSecret
	lineConfigs.LineChannelAccessToken = dataCreate.LineChannelAccessToken

	result, err := lineConfigsService.CreateLineConfig(lineConfigs)
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
		return
	}

	if result == "Line Channel Id is taken" {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, result))
		return
	}

	ctx.Response().Json(services.NewResMessageService().Json(http.StatusCreated, result))
	return
}

func (r *LineController) UpdateLineConfig(ctx myHttp.Context) {
	var dataUpdate = new(updateLineConfig)

	err := ctx.Request().Bind(&dataUpdate)
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusBadRequest, err))
		return
	}

	lineChannelId := ctx.Request().Query("lineChannelId")

	var dataLineConfig = new(models.LineConfigs)
	dataLineConfig.LineChannelName = dataUpdate.LineChannelName
	dataLineConfig.LineChannelAccessToken = dataUpdate.LineChannelAccessToken
	dataLineConfig.LineChannelSecret = dataUpdate.LineChannelSecret

	result, err := lineConfigsService.UpdateLineConfig(lineChannelId, *dataLineConfig)
	if err != nil {
		ctx.Response().Json(services.NewResMessageService().Json(http.StatusInternalServerError, err))
		return
	}

	ctx.Response().Json(services.NewResMessageService().Json(http.StatusOK, result))
	return
}

func handleUserLine(userLineId string, userLineType string) (*uint, error) {
	var user models.Users
	err := facades.Orm().Query().FindOrFail(&user, "user_line_id=?", userLineId)

	if err != nil {
		user.UserLineID = userLineId
		user.UserType = userLineType
		CreateUser := facades.Orm().Query().Create(&user)

		if CreateUser != nil {
			return nil, CreateUser
		}
	}

	return &user.Id, nil
}

func handleImageMessage(bot *linebot.Client, message *linebot.ImageMessage, userId *uint, LineConfigId uint) (*uint, error) {
	// Get image content
	content, err := bot.GetMessageContent(message.ID).Do()
	if err != nil {
		log.Println("Error retrieving image content:", err)
		return nil, err
	}
	defer content.Content.Close()

	// Save image to a file
	filename := fmt.Sprintf("%s.jpg", message.ID)
	fileDir := "storage/img"
	filePath := filepath.Join(fileDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating image file:", err)
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, content.Content)
	if err != nil {
		log.Println("Error saving image file:", err)
		return nil, err
	}

	var MessageImage models.MessagesReceivedImage
	MessageImage.ContentProvider = filePath
	MessageImage.MessageLineID = message.ID
	MessageImage.UserID = *userId
	MessageImage.LineConfigID = LineConfigId
	err = facades.Orm().Query().Create(&MessageImage)
	if err != nil {
		return nil, err
	}

	messageType := models.MessageTypes{MessageTypeSlug: "img"}
	err = facades.Orm().Query().Find(&messageType)
	if err != nil {
		// Handle error
		return nil, err
	}

	userMessageType := models.UserMessageTypes{}
	userMessageType.MessageId = MessageImage.Id
	userMessageType.UserID = MessageImage.UserID
	userMessageType.MessageTypeID = messageType.Id
	userMessageType.LineConfigID = LineConfigId
	err = facades.Orm().Query().Create(&userMessageType)
	if err != nil {
		return nil, err
	}

	return &userMessageType.Id, nil
}
func handleTextMessage(message *linebot.TextMessage, userId *uint, lineConfigId uint) (*uint, error) {
	// Get text content
	var MessageText models.MessagesReceivedText
	MessageText.MessageLineID = message.ID
	MessageText.MessageText = message.Text
	MessageText.UserID = *userId
	MessageText.LineConfigID = lineConfigId

	err := facades.Orm().Query().Create(&MessageText)
	if err != nil {
		return nil, err
	}

	messageType := models.MessageTypes{MessageTypeSlug: "text"}
	err = facades.Orm().Query().Find(&messageType)
	if err != nil {
		// Handle error
		return nil, err
	}

	userMessageType := models.UserMessageTypes{}
	userMessageType.MessageId = MessageText.Id
	userMessageType.UserID = MessageText.UserID
	userMessageType.MessageTypeID = messageType.Id
	userMessageType.LineConfigID = lineConfigId
	err = facades.Orm().Query().Create(&userMessageType)
	if err != nil {
		return nil, err
	}

	return &userMessageType.Id, nil
}

func handleStickerMessage(message *linebot.StickerMessage, userId *uint, lineConfigId uint) (*uint, error) {
	var MessageSticker models.MessagesReceivedSticker
	MessageSticker.MessageLineID = message.ID
	MessageSticker.StickerId = message.StickerID
	MessageSticker.StickerResourceType = string(message.StickerResourceType)
	MessageSticker.PackageId = message.PackageID
	MessageSticker.UserID = *userId
	MessageSticker.LineConfigID = lineConfigId
	err := facades.Orm().Query().Create(&MessageSticker)
	if err != nil {
		return nil, err
	}
	err = facades.Orm().Query().FindOrFail(&MessageSticker, "message_line_id=?", message.ID)
	if err != nil {
		return nil, err
	}

	messageType := models.MessageTypes{MessageTypeSlug: "sticker"}
	err = facades.Orm().Query().Find(&messageType)
	if err != nil {
		// Handle error
		return nil, err
	}

	userMessageType := models.UserMessageTypes{}
	userMessageType.MessageId = MessageSticker.Id
	userMessageType.UserID = MessageSticker.UserID
	userMessageType.MessageTypeID = messageType.Id
	userMessageType.LineConfigID = lineConfigId
	err = facades.Orm().Query().Create(&userMessageType)
	if err != nil {
		return nil, err
	}

	return &userMessageType.Id, nil
}

func handleLocationMessage(message *linebot.LocationMessage, userId *uint, lineConfigId uint) (*uint, error) {
	var MessageLocation models.MessagesReceivedLocation
	MessageLocation.MessageLineID = message.ID
	MessageLocation.UserID = *userId
	MessageLocation.Address = message.Address
	MessageLocation.Latitude = fmt.Sprintf("%f", message.Latitude)
	MessageLocation.Longitude = fmt.Sprintf("%f", message.Longitude)
	MessageLocation.LineConfigID = lineConfigId

	err := facades.Orm().Query().Create(&MessageLocation)
	if err != nil {
		return nil, err
	}
	err = facades.Orm().Query().FindOrFail(&MessageLocation, "message_line_id=?", message.ID)
	if err != nil {
		return nil, err
	}

	messageType := models.MessageTypes{MessageTypeSlug: "locations"}
	err = facades.Orm().Query().Find(&messageType)
	if err != nil {
		// Handle error
		return nil, err
	}

	userMessageType := models.UserMessageTypes{}
	userMessageType.MessageId = MessageLocation.Id
	userMessageType.UserID = MessageLocation.UserID
	userMessageType.MessageTypeID = messageType.Id
	userMessageType.LineConfigID = lineConfigId
	err = facades.Orm().Query().Create(&userMessageType)
	if err != nil {
		return nil, err
	}

	return &userMessageType.Id, nil
}

func handleAudioMessage(bot *linebot.Client, message *linebot.AudioMessage, userId *uint, lineConfigId uint) (*uint, error) {
	// Access the audio content using message.Content() function
	audioContent, err := bot.GetMessageContent(message.ID).Do()
	if err != nil {
		// Handle error
		return nil, err
	}

	// Create a new file to save the audio
	filename := fmt.Sprintf("%s.m4a", message.ID)
	fileDir := "storage/audio"
	filePath := filepath.Join(fileDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating audio file:", err)
		return nil, err
	}
	defer file.Close()

	// Write the audio content to the file
	_, err = io.Copy(file, audioContent.Content)
	if err != nil {
		// Handle error
		return nil, err
	}

	var MessageAudio models.MessagesReceivedAudio
	MessageAudio.MessageLineID = message.ID
	MessageAudio.Duration = strconv.Itoa(message.Duration)
	MessageAudio.ContentProvider = filePath
	MessageAudio.UserID = *userId
	MessageAudio.LineConfigID = lineConfigId
	facades.Orm().Query().Create(&MessageAudio)
	err = facades.Orm().Query().FindOrFail(&MessageAudio, "message_line_id=?", message.ID)
	if err != nil {
		// Handle error
		return nil, err
	}

	messageType := models.MessageTypes{MessageTypeSlug: "audio"}
	err = facades.Orm().Query().Find(&messageType)
	if err != nil {
		// Handle error
		return nil, err
	}
	userMessageType := models.UserMessageTypes{}
	userMessageType.MessageId = MessageAudio.Id
	userMessageType.UserID = MessageAudio.UserID
	userMessageType.MessageTypeID = messageType.Id
	userMessageType.LineConfigID = lineConfigId
	err = facades.Orm().Query().Create(&userMessageType)
	if err != nil {
		return nil, err
	}

	return &userMessageType.Id, nil
}

func handleRepliedMessage(bot *linebot.Client, replyToken string, typeMessage string, id uint, userId uint, lineConfigId uint) {
	messageTextReplied := "Thank you. Your message type: " + typeMessage
	userAdmin := models.UserAdmins{UserName: "systems"}
	err := facades.Orm().Query().Find(&userAdmin)
	if err != nil {
		return
	}
	var replyModel models.RepliesMessage
	replyModel.MessageText = messageTextReplied
	replyModel.UserAdminID = userAdmin.Id
	replyModel.UserID = userId
	replyModel.LineConfigID = lineConfigId
	replyModel.UserMessageTypesID = id

	err = facades.Orm().Query().Create(&replyModel)
	if err != nil {
		log.Println("Save log replying Error:", err)
		return
	}

	_, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(messageTextReplied)).Do()
	if err != nil {
		log.Println("Error replying:", err)
		return
	}
}

func (r *LineController) TestDb(ctx myHttp.Context) {
	var user models.Users
	facades.Orm().Query().First(&user)
	log.Print(user)
	// fmt.Print(err)
}
