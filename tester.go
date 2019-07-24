package main

import (
	"./src/GODP"
	"github.com/common-nighthawk/go-figure"
)

func main() {

	odpOBJ := GODP.ODP{}
	odpOBJ.BaseURL = "https://cloud.odp.capiot.com"
	odpOBJ.Username = "jaideep@capiot.com"
	odpOBJ.Password = "917875221075"
	odpOBJ.AppName = "Bajaj"
	odpOBJ.ServiceName = "odptester"
	odpOBJ.JWT = ""
	figure.NewFigure("Go O.D.P !", "", true).Print()

	println("\nodp object initialized\n")
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

	testBodyMap["name"] = "jaideep_updated"
	println("Update Item started firing")
	resp4, status4 := odpOBJ.UpdateItem(testBodyMap, resp3["_id"].(string))
	println("Update Item fired with response code:", status4,
		" and created updated ODP object with id:", resp4["_id"].(string),
		"with new val ", resp4["name"].(string))

	if status1 != 200 {
		panic("Get Items Failed")
	}

	if status2 != 200 {
		panic("Get Item Failed")
	}

	if status3 != 200 {
		panic("Create Items Failed")
	}

}
