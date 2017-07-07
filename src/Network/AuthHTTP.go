package Network

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"Authorization/AuthorizationNetwork"
)



func Register(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		setServerError(w)
		return
	}

	//Только так работает :(
	var decoder = json.NewDecoder(strings.NewReader(string(body)))

	var params AuthorizationNetwork.RegisterInfo
	err = decoder.Decode(&params)

	if err != nil && err != io.EOF {
		setServerError(w)
		return
	}

	err = authExecutor.Register(params)

	if err != nil {
		setServerError(w)
		return
	}

	var successRes, err2 = getSimpleResult(true)
	
	if err2 != nil {
		setServerError(w)
		return
	}
	
	json.NewEncoder(w).Encode(successRes)
}




func getSimpleResult(success bool) ([]byte, error) {
	
	var result = make(map[string]interface{})
	result["success"] = success
	
	return json.Marshal(result)
}