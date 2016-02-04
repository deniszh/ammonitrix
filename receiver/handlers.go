package receiver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/backends/elastic"
	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

func (r *Receiver) handleData(w http.ResponseWriter, req *http.Request) {
	log.Printf("[DEBUG] Received datagram")
	if req.Method != "PUT" {
		http.Error(w, "Unsupported method", 405)
		return
	}
	var valid bool
	var datagram config.ElasticData
	if valid, datagram = r.validateDataRequest(req.Body); valid != true {
		http.Error(w, "Invalid request received", 500)
		return
	}
	log.Printf("[DEBUG] Received a valid datagram: %s", datagram)

	dResp, dErr := r.Elastic.StoreDatagram(datagram)
	if dErr != nil {
		http.Error(w, "Failed to store datagram", 500)
	}

	w.WriteHeader(dResp.StatusCode)
}

func (r *Receiver) handleAPI(w http.ResponseWriter, req *http.Request) {
	log.Printf("[DEBUG] Received API call")

	if req.Method == "GET" {
		el, err := elastic.NewElastic(r.Config)
		if err != nil {
			http.Error(w, "Error initializing Elastic", 500)
			return
		}

		r.Elastic = el
		body, _ := r.Elastic.SearchAll()
		fmt.Fprintf(w, "%s", body)
		return
	}

	if req.Method != "POST" { // ie PUT, DELETE, etc.
		http.Error(w, "Unsupported method", 405)
		return
	}

	decoder := json.NewDecoder(req.Body)
	var j config.APISearch
	err := decoder.Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println(j)
}
