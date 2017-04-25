// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

//PetDB :
//var PetDB *Pets

func main() {
	var err error
        //PetDB = NewPets()
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {	
			case *linebot.TextMessage:
				/*var pet *Pet*/
				log.Println(message.Text)
				inText := strings.ToLower(message.Text)
				if strings.Contains(inText, "三色線") || strings.Contains(inText, "三色線接法") || strings.Contains(inText, "三色線怎麼接") || strings.Contains(inText, "三色線接電視") || strings.Contains(inText, "三條線") || strings.Contains(inText, "三條線接法") || strings.Contains(inText, "三條線怎麼接") || strings.Contains(inText, "三條線接電視") || strings.Contains(inText, "AV線") || strings.Contains(inText, "AV線接法") || strings.Contains(inText, "AV線怎麼接") || strings.Contains(inText, "AV線接電視") || strings.Contains(inText, "AV端子") || strings.Contains(inText, "AV端子接法") || strings.Contains(inText, "AV端子怎麼接") || strings.Contains(inText, "AV端子接電視") || strings.Contains(inText, "電視接三色線") || strings.Contains(inText, "電視接三條線") || strings.Contains(inText, "電視接AV線") || strings.Contains(inText, "電視接AV端子"){
					out := fmt.Sprintf(`接法為:黃接黃(影像);白接白(左聲道);紅接紅(右聲道)
                                                            參考影片:https://www.youtube.com/watch?v=j80aHfxMRXE`)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
					}
				}else if strings.Contains(inText, "客服專線") || strings.Contains(inText, "客服") || strings.Contains(inText, "連接客服") || strings.Contains(inText, "打給客服") || strings.Contains(inText, "客服人員") || strings.Contains(inText, "電話"){
					out := fmt.Sprintf("04-26881407")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
					}
				}else if strings.Contains(inText, "你好") || strings.Contains(inText, "hello") || strings.Contains(inText, "Hello") || strings.Contains(inText, "HELLO") || strings.Contains(inText, "你好啊") || strings.Contains(inText, "你好呀"){
					out := fmt.Sprintf("你好")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
					}
				}else if strings.Contains(inText, "接電視沒聲音") || strings.Contains(inText, "接電視沒影像") || strings.Contains(inText, "沒聲音") || strings.Contains(inText, "沒影像"){
					out := fmt.Sprintf("請問是用什麼線?")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
					}
				}else if strings.Contains(inText, "HDMI") || strings.Contains(inText, "hdmi") || strings.Contains(inText, "HDMI線") || strings.Contains(inText, "hdmi線") || strings.Contains(inText, "高清線"){
					out := fmt.Sprintf("參考影片:https://www.youtube.com/watch?v=AbctqNEfkJw")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
					}
				}else if strings.Contains(inText, "作者") || strings.Contains(inText, "開發人員") || strings.Contains(inText, "Creater") || strings.Contains(inText, "CREATER") || strings.Contains(inText, "創作者"){
					out := fmt.Sprintf("Made by 蔡侑憬")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
					}
				}else if strings.Contains(inText, "網站") || strings.Contains(inText, "官方網站") || strings.Contains(inText, "韻全官方網站") || strings.Contains(inText, "韻全官網") || strings.Contains(inText, "官網") || strings.Contains(inText, "Official website") || strings.Contains(inText, "Official") || strings.Contains(inText, "official") || strings.Contains(inText, "official website"){
					out := fmt.Sprintf("傳送門:http://yunchuan1.weebly.com/")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
					}
				}else if strings.Contains(inText, "工程委託") || strings.Contains(inText, "工程") || strings.Contains(inText, "委託") || strings.Contains(inText, "委託工程") || strings.Contains(inText, "engineering") || strings.Contains(inText, "Engineering"){
					out := fmt.Sprintf("傳送門:http://yunchuan1.weebly.com/22996353513492021934.html")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
					}
				}
				/*if strings.Contains(inText, "狗") || strings.Contains(inText, "dog") {
					pet = PetDB.GetNextDog()
				}else if strings.Contains(inText, "貓") || strings.Contains(inText, "cat") {
					pet = PetDB.GetNextCat()
				}

				if pet == nil {
					pet = PetDB.GetNextPet()
				}*/
				
				/*out := fmt.Sprintf("您好，目前的動物：名為%s, 所在地為:%s, 敘述: %s 電話為:%s 圖片網址在: %s", pet.Name, pet.Resettlement, pet.Note, pet.Phone, pet.ImageName)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(out)).Do(); err != nil {
					log.Print(err)
				}

				log.Println("Img:", pet.ImageName)

				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(pet.ImageName, pet.ImageName)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
	 // Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+" 早安 : ) 快樂星期天呦! Made by 蔡侑憬")).Do(); err != nil {
					log.Print(err)
				}*/
			}
		}
	}
}

//  message.Text  
