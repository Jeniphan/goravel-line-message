package line

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goravel/app/models"
	"io/ioutil"
	"log"
	"net/http"

	myHttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type LineMessage struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Timestamp  int64  `json:"timestamp"`
		Source     struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		Message struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

type ReplyMessage struct {
	ReplyToken string `json:"replyToken"`
	Messages   []Text `json:"messages"`
}
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

var channalToken = "EcaHmTA6D+qykg0l9cB+KpgnHQw2OdEi/KxQZsu1uMUqzYqawwWMo9UYY0PXwY53MdX25XDixJNwDyxHWHTfHutpdj5E0aq2lgGrdcP0O34wOzwVSilRq64Y9GrOmlh+BI5AUipzvX2pzNo+1wjo3QdB04t89/1O/w1cDnyilFU="

type LineController struct {
	//Dependent services
}

func NewLineController() *LineController {
	return &LineController{
		//Inject services
	}
}

func (r *LineController) MessageLine(ctx myHttp.Context) {
	var Line LineMessage
	_ = json.NewDecoder(ctx.Request().Origin().Body).Decode(&Line)

	fmt.Println("UserId:", Line.Events[0].Source.UserID, " : Message", Line.Events[0].Message.Text)

	text := Text{
		Type: "text",
		Text: "ข้อความเข้ามา : " + Line.Events[0].Message.Text + " ยินดีต้อนรับ : " + Line.Events[0].Source.UserID,
	}

	message := ReplyMessage{
		ReplyToken: Line.Events[0].ReplyToken,
		Messages: []Text{
			text,
		},
	}

	go replyMessageLine(message)

}

func replyMessageLine(Message ReplyMessage) error {
	value, _ := json.Marshal(Message)
	url := "https://api.line.me/v2/bot/message/reply"

	var jsonStr = []byte(value)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+channalToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return err
}

func (r *LineController) TestDb(ctx myHttp.Context) {
	var user models.Users
	facades.Orm().Query().First(&user)
	log.Print(user)
	// fmt.Print(err)
}
