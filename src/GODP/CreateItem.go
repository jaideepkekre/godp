package GODP

import "encoding/json"
import "bytes"
import "io/ioutil"

//Create an Item in ODP
func (odpOBJ *ODP) CreateItem(input map[string]interface{}) (map[string]interface{}, int) {

	trailURL := "/api/c/" + odpOBJ.AppName + "/" + odpOBJ.ServiceName

	body, err_marshalling := json.Marshal(input)
	logerror(err_marshalling)
	req := odpOBJ.createRequest("POST", trailURL, bytes.NewBuffer(body))
	resp, err_req := netClient.Do(req)
	logerror(err_req)
	if resp.StatusCode ==403 {
		odpOBJ.Login()
		resp = nil
		err_req = nil
		resp, err_req = netClient.Do(req)
	}
	bytes, err_bytes := ioutil.ReadAll(resp.Body)
	logerror(err_bytes)
	resMap := processPureResponse(bytes)
	return resMap, resp.StatusCode

}

