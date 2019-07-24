package GODP

import "io/ioutil"

func (odpOBJ *ODP) Get(itemid string) (map[string]interface{}, int) {

	//STEP 1
	trailURL := "/api/c/" + odpOBJ.AppName + "/" + odpOBJ.ServiceName + "/" + itemid
	req := odpOBJ.createRequest("GET", trailURL, nil)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	//STEP 2
	resp, err_req := netClient.Do(req)
	logerror(err_req)

	//STEP 3 OPTIONAL
	if resp.StatusCode == 403 {
		odpOBJ.Login()
		resp = nil
		err_req = nil
		resp, err_req = netClient.Do(req)
	}

	//STEP 4
	bytes, err_bytes := ioutil.ReadAll(resp.Body)
	logerror(err_bytes)
	respMap := processPureJSONResponse(bytes)
	return respMap, resp.StatusCode

}

