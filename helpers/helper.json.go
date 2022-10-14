package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/edot-interview/models"
)

func Strigify(payload interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

func Parse(payload []byte) Response {
	var jsonResponse Response
	err := json.Unmarshal(payload, &jsonResponse)
	fmt.Println(jsonResponse)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return jsonResponse
}
func ParseProduct(payload []byte) models.Product {
	var jsonResponse models.Product
	err := json.Unmarshal(payload, &jsonResponse)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return jsonResponse
}
