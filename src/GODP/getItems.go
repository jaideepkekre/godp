package GODP

import "io/ioutil"

// import "net/http"

func (odpOBJ *ODP) GetItems(queryParams map[string]string) ([]interface{}, int) {
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
	println(req.URL.String())
	resp, err_req := netClient.Do(req)
	logerror(err_req)
	if resp.StatusCode == 403 {
		odpOBJ.Login()
		resp = nil
		err_req = nil
		resp, err_req = netClient.Do(req)
	}
	bytes, err_bytes := ioutil.ReadAll(resp.Body)
	logerror(err_bytes)
	respMap := processArrayResponse(bytes)
	return respMap, resp.StatusCode

}

