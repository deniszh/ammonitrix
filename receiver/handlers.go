package receiver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

type api_search struct {
	State                 string
	Current_state_time    string
	Current_state_updates string
	Quiet                 string
}

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
		fmt.Fprintf(w, "insert all checks here\n")
		return
	}

	if req.Method != "POST" {
		http.Error(w, "Unsupported method", 405)
		return
	}

	decoder := json.NewDecoder(req.Body)
	var j api_search
	err := decoder.Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println(j)
}
