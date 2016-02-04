package receiver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

func (r *Receiver) validateDataRequest(body io.Reader) (bool, config.ElasticData) {

	var elasticData config.ElasticData
	var elasticMeta config.ElasticMetadata

	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	body_string := buf.String()

	log.Println("[DEBUG] body:", body_string)

	log.Println("[DEBUG] ElasticData before decoding:", elasticData)

	decoder := json.NewDecoder(buf)

	err := decoder.Decode(&elasticData)
	if err != nil {
		log.Println("[ERROR] Failed unmarshalling ElasticData:", err)
		return false, elasticData
	}

	log.Println("[DEBUG] ElasticData after decoding:", elasticData)

	decoder = json.NewDecoder(bytes.NewBufferString(body_string))
	err = decoder.Decode(&elasticMeta)
	if err != nil {
		log.Println("[ERROR] Failed unmarshalling ElasticMeta:", err)
		return false, elasticData
	}

	log.Println("[DEBUG] ElasticMeta after decoding:", elasticMeta)

	var w http.ResponseWriter
	resp, err := r.Elastic.StoreRegistration(elasticMeta)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to store register, ElasticSearch response code: %d", resp.StatusCode), 500)
	}

	return true, elasticData
}
