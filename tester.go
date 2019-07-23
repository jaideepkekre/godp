package main

import (
	"./src/GODP"
)

func main() {
	// trailurl := "/api/c/" + appName + "/" + serviceName

	odpOBJ := GODP.ODP{}
	odpOBJ.BaseURL = "https://cloud.odp.capiot.com"
	odpOBJ.Username = "jaideep@capiot.com"
	odpOBJ.Password = "917875221075"
	odpOBJ.AppName = "Bajaj"
	odpOBJ.ServiceName = "odptester"
	odpOBJ.JWT=""
	var r =make(map[string]string)
	r["count"] = "5"
	r["filter"] = "{\"name\":\"Test\"}"
	r["sort"] = "id"
	r["page"] = "45"

	resp,_:=odpOBJ.GetItems(r)
	print(len(resp))

	resp2,_ := odpOBJ.Get("ODP1144")
	print(resp2["name"].(string))

	resp3,status3 :=odpOBJ.CreateItem(resp2)
	print(resp3["name"].(string))
	print(status3)


}
