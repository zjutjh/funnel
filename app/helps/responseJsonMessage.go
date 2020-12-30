package helps

import (
	"encoding/json"
	"log"
)
import "funnel/app/errors"

type ResponseJsonMessage struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponseJson(d interface{}) []byte {
	jsonRes, err := json.Marshal(ResponseJsonMessage{Code: 200, Message: "OK", Data: d})
	if err != nil {
		log.Fatal(err.Error())
	}
	return jsonRes
}

func FailResponseJson(error errors.ResponseError, d interface{}) []byte {
	jsonRes, _ := json.Marshal(ResponseJsonMessage{Code: error.Code, Message: error.Message, Data: d})
	return jsonRes
}
