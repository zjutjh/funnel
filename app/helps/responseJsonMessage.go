package helps

import "encoding/json"
import "funnel/app/errors"

type ResponseJsonMessage struct {
	code    int
	message string
	data    interface{}
}

func SuccessResponseJson(d interface{}) []byte {
	jsonRes, _ := json.Marshal(ResponseJsonMessage{code: 200, message: "OK", data: d})
	return jsonRes
}

func FailResponseJson(error errors.ResponseError, d interface{}) []byte {
	jsonRes, _ := json.Marshal(ResponseJsonMessage{code: error.Code, message: error.Message, data: d})
	return jsonRes
}
