package GODP

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func processPureJSONResponse(resp_body []byte) map[string]interface{} {
	var respBodyInterface interface{}
	err_unmarshal := json.Unmarshal(resp_body, &respBodyInterface)
	if err_unmarshal != nil {
		log.Fatalln(err_unmarshal)
	}
	// 3
	respBodyMap := respBodyInterface.(map[string]interface{})
	return respBodyMap

}

func processJSONArrayResponse(resp_body []byte) []interface{} {
	var respBodyInterface interface{}
	err_unmarshal := json.Unmarshal(resp_body, &respBodyInterface)
	if err_unmarshal != nil {
		log.Fatalln(err_unmarshal)
	}
	// 3
	respBodyMap := respBodyInterface.([]interface{})
	return respBodyMap

}

func logerror(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func GetMapfromInterface(input interface{}) map[string]interface{} {
	return (input.(map[string]interface{}))

}

func (odpOBJ *ODP) createRequest(typ string, trailURL string, payload *bytes.Buffer) *http.Request {
	url := odpOBJ.BaseURL + trailURL
	var req *http.Request
	switch typ {
	case "GET":
		req, _ = http.NewRequest("GET", url, nil)
	case "POST":
		req, _ = http.NewRequest("POST", url, payload)
	default:

	}
	if odpOBJ.JWT == "" {
		odpOBJ.Login()
	}
	req.Header.Set("Authorization", "JWT "+odpOBJ.JWT)
	req.Header.Set("Content-Type", "application/json")
	return req

}

