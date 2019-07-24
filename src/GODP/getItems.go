package GODP

import "io/ioutil"

func (odpOBJ *ODP) GetItems(queryParams map[string]string) ([]interface{}, int) {

	//Step 1
	trailURL := "/api/c/" + odpOBJ.AppName + "/" + odpOBJ.ServiceName
	req := odpOBJ.createRequest("GET", trailURL, nil)
	q := req.URL.Query()
	if filtr, ok := queryParams["filter"]; ok {
		q.Add("filter", filtr)

	}
	if srt, ok := queryParams["sort"]; ok {
		q.Add("sort", srt)

	}

	if pg, ok := queryParams["page"]; ok {
		q.Add("page", pg)

	}

	if cnt, ok := queryParams["count"]; ok {
		q.Add("count", cnt)
	}

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
	respMap := processJSONArrayResponse(bytes)
	return respMap, resp.StatusCode

}

