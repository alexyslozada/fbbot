package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"bytes"
	"io/ioutil"
)

type CommonFormat struct {
	Object string `json:"object"`
	Entry  []struct {
		ID        string `json:"id"`
		Time      int64  `json:"time"`
		Messaging []struct {
			Sender struct {
				ID string `json:"id"`
			} `json:"sender"`
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Timestamp int64 `json:"timestamp"`
			Message   struct {
				Mid  string `json:"mid"`
				Seq  int64  `json:"seq"`
				Text string `json:"text"`
			} `json:"message"`
		} `json:"messaging"`
	} `json:"entry"`
}

type ResponseMessage struct {
	Recipient struct {
		ID string `json:"id,omitempty"`
	} `json:"recipient,omitempty"`
	Message struct {
		Text       string `json:"text,omitempty"`
		/*
		Attachment struct {
			Type    string `json:"type,omitempty"`
			Payload struct {
				Url          string `json:"url,omitempty"`
				TemplateType string `json:"template_type,omitempty"`
				Elements     []struct {
					Title    string `json:"title,omitempty"`
					Subtitle string `json:"subtitle,omitempty"`
					ItemUrl  string `json:"item_url,omitempty"`
					ImageUrl string `json:"image_url,omitempty"`
					Buttons  []struct {
						Type  string `json:"type,omitempty"`
						Url   string `json:"url,omitempty"`
						Title string `json:"title,omitempty"`
					} `json:"buttons,omitempty"`
				} `json:"elements,omitempty"`
			} `json:"payload,omitempty"`
		} `json:"attachment,omitempty"`
		*/
	} `json:"message,omitempty"`
}

func main() {
	fmt.Println("Hola")
	http.HandleFunc("/", sayhello)
	http.HandleFunc("/webhook", webhook)
	log.Println(http.ListenAndServeTLS(":8181", "/etc/letsencrypt/archive/app.golang-es.com/fullchain1.pem", "/etc/letsencrypt/archive/app.golang-es.com/privkey1.pem", nil))
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo"))
}

func webhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Han hecho una petición")
	if r.Method == http.MethodGet {
		log.Println("Método: GET")
		vt := r.URL.Query().Get("hub.verify_token")
		if vt == "test_bot_token" {
			w.Write([]byte(r.URL.Query().Get("hub.challenge")))
			return
		} else {
			log.Print("Token no valido")
			w.Write([]byte("Token no valido"))
			return
		}
	}

	if r.Method == http.MethodPost {
		log.Println("Método: POST")
		cf := CommonFormat{}
		err := json.NewDecoder(r.Body).Decode(&cf)
		if err != nil {
			log.Print(err)
			return
		}

		if cf.Object == "page" {
			for _, entry := range cf.Entry {
				for _, messaging := range entry.Messaging {
					if messaging.Message.Text != "" {
						rm := ResponseMessage{
							Recipient: struct {
								ID string `json:"id,omitempty"`
							}{
								ID: messaging.Sender.ID,
							},
							Message: struct {
								Text string `json:"text,omitempty"`
							}{
								Text: "Hola amigos!!!",
							},
						}

						j, err := json.Marshal(rm)
						if err != nil {
							log.Print("Error al crear el json")
							return
						}
						// w.Write(j)
						log.Println(string(j))
						req, err := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token=EAAKf9ZAASAsoBAGFgvHotpjA4kpkkgPhGOIuJcaZBDlZCE1P9QiBnJbPrMjOouvn7u50Nvtb9ikb3p7WPNerxZCJZBYFlVTeCvCLZCNVxsE3uc4IyIS4e1uBMJICX0jS9w3Cgk4SvPvnrDU9q6yuyxqyjfXMJgkdToPs3lECnBhAZDZD", bytes.NewBuffer(j))
						if err != nil {
							log.Print(err)
							return
						}
						req.Header.Set("Content-Type", "application/json")

						client := &http.Client{}
						resp, err := client.Do(req)
						if err != nil {
							log.Print(err)
							return
						}
						defer resp.Body.Close()

						log.Println("response Status:", resp.Status)
						log.Println("response Headers:", resp.Header)
						body, _ := ioutil.ReadAll(resp.Body)
						log.Println("response Body:", string(body))
						return
					}
				}
			}
		}
	}
}
