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
	odpOBJ.JWT = ""

	println("odp object initialized\n")
	var r = make(map[string]string)
	r["count"] = "5"
	r["filter"] = "{\"name\":\"Test\"}"
	r["sort"] = "id"
	r["page"] = "2"

	println("Get Items started firing")
	resp, status1 := odpOBJ.GetItems(r)
	println("Get Items fired with response code:", status1, "and got ", len(resp), " items\n")

	println("Get one started firing")
	resp2, status2 := odpOBJ.Get("ODP1144")
	println("Get one fired with response code:", status2, "and got:", resp2["name"].(string), "\n")

	testBodyMap := make(map[string]interface{})
	testBodyMap["name"] = "jaideep"

	println("Create Item started firing")
	resp3, status3 := odpOBJ.CreateItem(testBodyMap)
	println("Create Item fired with response code:", status3, " and created ODP object with id:", resp3["_id"].(string))

}
