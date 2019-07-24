package GODP

import "encoding/json"
import "bytes"
import "io/ioutil"

//Create an Item in ODP
func (odpOBJ *ODP) CreateItem(input map[string]interface{}) (map[string]interface{}, int) {
	//STEP 1
	trailURL := "/api/c/" + odpOBJ.AppName + "/" + odpOBJ.ServiceName

	body, err_marshalling := json.Marshal(input)
	logerror(err_marshalling)

	//Step2
	req := odpOBJ.createRequest("POST", trailURL, bytes.NewBuffer(body))
	resp, err_req := netClient.Do(req)
	logerror(err_req)

	//Step 3 OPTIONAL
	if resp.StatusCode == 403 {
		odpOBJ.Login()
		resp = nil
		err_req = nil
		resp, err_req = netClient.Do(req)
	}

	//STEP 4
	bytes, err_bytes := ioutil.ReadAll(resp.Body)
	logerror(err_bytes)
	resMap := processPureJSONResponse(bytes)
	return resMap, resp.StatusCode

}

