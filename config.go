package main

import (
	"io/ioutil"
	"github.com/labstack/gommon/log"
	"encoding/json"
)

// c es la variable que tiene toda la información
// del archivo de configuración
var c Config

// Config estructura de los datos variables de la aplicación
type Config struct {
	Port    string `json:"port"`
	CertPem string `json:"cert_pem"`
	PrivPem string `json:"priv_pem"`
	MyToken string `json:"my_token"`
	FbToken string `json:"fb_token"`
	FbUrl   string `json:"fb_url"`
}

// getConfiguration Lee la información
// del archivo config.json
func getConfiguration() {
	log.Print("Leyendo la información del config.json...")
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("Se encontró un error al leer el archivo config.json: %v", err)
	}

	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Fatalf("Error al convertir el json a estructura: %v", err)
	}

	log.Print("Lectura del archivo finalizada...")
}

