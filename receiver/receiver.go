package receiver

import (
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/backends/elastic"
	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

type Receiver struct {
	Config  *config.Config
	Elastic *elastic.Elastic
}

func NewReceiver(config *config.Config) (*Receiver, error) {
	r := &Receiver{
		Config: config,
	}

	return r, nil
}

var quit = make(chan error)

//StartListener reads elasticSearch to load existing metadata
//var registration map[string]config.ElasticMetadata

func (r *Receiver) StartListener(map[string]config.ElasticMetadata) error {
	http.HandleFunc("/data", r.handleData)
	go func() {
		err := http.ListenAndServe(r.Config.Listen.Port, nil)
		if err != nil {
			quit <- err
		}
	}()

	// wait for shutdown signal
	err := <-quit

	// trigger graceful shutdown
	log.Print("[INFO] Down: ", err)

	return nil
}

func (r *Receiver) ConnectElastic() (map[string]config.ElasticMetadata, error) {
	el, err := elastic.NewElastic(r.Config)
	if err != nil {
		return nil, err
	}

	r.Elastic = el
	var registration, _ = r.Elastic.LoadRegistration()

	return registration, err
}
