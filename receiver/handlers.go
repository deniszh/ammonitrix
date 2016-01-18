package receiver

import (
	"log"
	"net/http"

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
