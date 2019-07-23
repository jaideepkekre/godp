package GODP

import "encoding/json"
import "bytes"
import "io/ioutil"
import "log"

// Login to ODP
func (odpOBJ *ODP) Login() {

	// Set old JWT as empty
	odpOBJ.JWT = ""
	loginURL := odpOBJ.BaseURL + "/api/a/rbac/login"

	// Create body for login req
	body, _ := json.Marshal(map[string]interface{}{
		"username": odpOBJ.Username,
		"password": odpOBJ.Password,
	})

	// make login req
	resp, err_post := netClient.Post(loginURL, "application/json", bytes.NewBuffer(body))
	if err_post != nil {
		log.Fatalln(err_post)
	}
	
	if resp.StatusCode == 200 {
		defer resp.Body.Close()

		// net --> bytes --> interface{} --> map[]
		// 1
		resp_body, err_read := ioutil.ReadAll(resp.Body)
		if err_read != nil {
			log.Fatalln(err_read)
		}
		// 2
		var respBodyInterface interface{}
		err_unmarshal := json.Unmarshal(resp_body, &respBodyInterface)
		if err_unmarshal != nil {
			log.Fatalln(err_unmarshal)
		}
		// 3
		respBodyMap := respBodyInterface.(map[string]interface{})
		odpOBJ.JWT = respBodyMap["token"].(string)
	} 	else {
		log.Fatalln("Could not Login!, check password")
	}

}

