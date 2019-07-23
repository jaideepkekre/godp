package GODP
import "io/ioutil"

// import "net/http"




func (odpOBJ *ODP) Get(itemid string) (map[string]interface{},int) {
	trailURL := "/api/c/" + odpOBJ.AppName + "/" + odpOBJ.ServiceName + "/" + itemid
	req := odpOBJ.createRequest("GET",trailURL,nil)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	println(req.URL.String())	 
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
	respMap := processPureResponse(bytes)
	return respMap,resp.StatusCode
	
}

