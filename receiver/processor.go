package receiver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

func (r *Receiver) validateDataRequest(body io.Reader) (bool, config.ElasticData) {
	decoder := json.NewDecoder(body)
	var d config.ElasticData
	var m config.ElasticMetadata
	err := decoder.Decode(&d)
	if err != nil {
		log.Println(err)
		return false, d
	}

	//TODO: check if need to register
	var w http.ResponseWriter
	resp, err := r.Elastic.StoreRegistration(m)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to store register, ElasticSearch response code: %d", resp.StatusCode), 500)
	}

	return true, d
}
