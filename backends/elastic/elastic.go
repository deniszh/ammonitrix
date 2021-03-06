package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

//Elastic Config for connecting to elastic
type Elastic struct {
	Config *config.Config
}

func NewElastic(config *config.Config) (*Elastic, error) {
	e := &Elastic{
		Config: config,
	}

	return e, nil
}

//StoreDatagram stores data
func (e *Elastic) StoreDatagram(elasticData config.ElasticData) (*http.Response, error) {
	url := fmt.Sprintf("http://%s%s/%s/event", e.Config.Elastic.Host, e.Config.Elastic.Port, e.Config.Elastic.IndexName)

	b, err := json.Marshal(elasticData)
	if err != nil {
		log.Println("[ERROR] Couldn't marshal datagram into JSON")
		return nil, err
	}

	log.Println("[DEBUG] Marshalled elasticData:", string(b))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Println("[ERROR] Something went wrong with http.NewRequest")
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}

/*LoadRegistration data on startup and return it as a map
FIXME: Function not completed
*/
func (e *Elastic) LoadRegistration() (map[string]config.ElasticMetadata, error) {
	log.Println("LOADING REGISTRATION")
	//FIXME: We want to scan for all docs, so this query is probably wrong.
	url := fmt.Sprintf("http://%s%s/%s/_search/?size=1000", e.Config.Elastic.Host, e.Config.Elastic.Port, e.Config.Elastic.MetaDataIndex)
	r, err := http.Get(url)
	if err != nil || r.StatusCode >= 400 {
		log.Println("[ERROR] Couldn't load existing metadata")
		return nil, err
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("[DEBUG] Got registration query response of:", string(body))
	var data config.Datagram
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Println("[ERROR] Could not unmarshal JSON")
	}

	return nil, nil
}

func (e *Elastic) StoreRegistration(elasticMeta config.ElasticMetadata) (*http.Response, error) {
	url := fmt.Sprintf("http://%s%s/%s/register", e.Config.Elastic.Host, e.Config.Elastic.Port, e.Config.Elastic.MetaDataIndex)

	log.Println("Storing registration")
	//Keep only metadata
	b, err := json.Marshal(elasticMeta)
	if err != nil {
		log.Println("[ERROR] Couldn't marshal datagram into JSON")
		return nil, err
	}

	log.Println("[DEBUG] Marshalled metadata:", string(b))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Println("[ERROR] Something went wrong with http.NewRequest")
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return response, nil
}

func (e *Elastic) SearchAll() ([]byte, error) {
	url := fmt.Sprintf("http://%s%s/%s/_search/", e.Config.Elastic.Host, e.Config.Elastic.Port, e.Config.Elastic.IndexName)
	log.Printf("[DEBUG] Searching in ES (%s)\n", url)
	r, err := http.Get(url)
	if err != nil || r.StatusCode >= 400 {
		log.Println("[ERROR] Couldn't search ES")
		return nil, err
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	return body, nil
}

func (e *Elastic) SearchKeyword(key string, value string) ([]byte, error) {
	url := fmt.Sprintf("http://%s%s/%s/_search?q=%s:%s", e.Config.Elastic.Host, e.Config.Elastic.Port, e.Config.Elastic.IndexName, key, value)
	log.Printf("[DEBUG] Searching in ES (%s)\n", url)
	r, err := http.Get(url)
	if err != nil || r.StatusCode >= 400 {
		log.Println("[ERROR] Couldn't search ES")
		return nil, err
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	return body, nil
}
