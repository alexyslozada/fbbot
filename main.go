package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"bytes"
	"io/ioutil"
	"github.com/alexyslozada/fbbot/model"
)

func main() {
	fmt.Println("Hola")

	// Se lee la información del config.json
	getConfiguration()

	http.HandleFunc("/", sayhello)
	http.HandleFunc("/webhook", webhook)
	log.Println(http.ListenAndServeTLS(c.Port, c.CertPem, c.PrivPem, nil))
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo"))
}

func webhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Han hecho una petición")
	if r.Method == http.MethodGet {
		log.Println("Método: GET")
		vt := r.URL.Query().Get("hub.verify_token")
		if vt == c.MyToken {
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
		cf := model.CommonFormat{}
		buf, _ := ioutil.ReadAll(r.Body)
		allData := ioutil.NopCloser(bytes.NewBuffer(buf))
		r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		leer := new(bytes.Buffer)
		leer.ReadFrom(allData)
		log.Printf("***Lo recibido es: %v", leer.String())
		err := json.NewDecoder(r.Body).Decode(&cf)
		if err != nil {
			log.Print(err)
			return
		}

		if cf.Object == "page" {
			for _, entry := range cf.Entry {
				for _, messaging := range entry.Messaging {
					if messaging.Message.Text != "" {
						rm := model.ResponseMessage{}
						if messaging.Message.Text == "llamar" {
							rm = model.ResponseMessage{
								Recipient: model.Recipient{
									ID: messaging.Sender.ID,
								},
								MessageContent: model.MessageContent{
									// Text: "Hola amigos!!!",
									Attachment: &model.Attachment{
										Type: "template",
										Payload: model.Payload{
											TemplateType: "generic",
											Elements: []model.Elements{
												model.Elements{
													Title: "Bienvenido amigo",
													ImageUrl: "https://lorempixel.com/190/100",
													ItemUrl: "https://app.golang-es.com",
													Subtitle: "Deseas llamarnos?",
													Buttons: []model.Buttons{
														model.Buttons{
															Type: "phone_number",
															Title: "Contáctanos al call center",
															Payload: "+573165036245",
														},
													},
												},
											},
										},
									},
								},
							}
						} else {
							rm = model.ResponseMessage{
								Recipient: model.Recipient{
									ID: messaging.Sender.ID,
								},
								MessageContent: model.MessageContent{
									Text: "Hola amigos!!!",
									Attachment: nil,
								},
							}
						}

						j, err := json.Marshal(rm)
						if err != nil {
							log.Print("Error al crear el json")
							return
						}
						// w.Write(j)
						log.Println(string(j))
						fu := fmt.Sprintf("%s?access_token=%s", c.FbUrl, c.FbToken)
						req, err := http.NewRequest("POST", fu, bytes.NewBuffer(j))
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
