package remote

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func convertToJSON(dest interface{}, r *http.Response) error {
	dataBytes, _ := ioutil.ReadAll(r.Body)
	jsonConverterErr := json.Unmarshal([]byte(dataBytes), dest)

	if jsonConverterErr != nil {
		return jsonConverterErr
	}

	return nil
}
