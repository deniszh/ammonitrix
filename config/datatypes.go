package config

import "encoding/json"

type ElasticMetadata struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func (o *ElasticMetadata) UnmarshalJSON(data []byte) error {
	var inter interface{}
	if err := json.Unmarshal(data, &inter); err != nil {
		return err
	}
	m := inter.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "name":
			o.Name = v.(string)
		case "tags":
			o.Tags = v.([]string)
		}
	}
	return nil
}

type ElasticData struct {
	Name      string      `json:"name"`
	CheckData interface{} `json:"check_data"`
}

func (o *ElasticData) UnmarshalJSON(data []byte) error {
	var inter interface{}
	if err := json.Unmarshal(data, &inter); err != nil {
		return err
	}
	m := inter.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "name":
			o.Name = v.(string)
		case "CheckData":
			o.CheckData = v.(map[string]interface{})
		}
	}
	return nil
}

type Datagram struct {
	Data     ElasticData     `json:data`
	Metadata ElasticMetadata `json:metadata`
}

func (o *Datagram) UnmarshalJSON(data []byte) error {
	var inter interface{}
	if err := json.Unmarshal(data, &inter); err != nil {
		return err
	}
	m := inter.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "data":
			o.Data = v.(ElasticData)
		case "metadata":
			o.Metadata = v.(ElasticMetadata)
		}
	}
	return nil
}

/*UnmarshalJSON for ElasticResponseMeta struct
Sample JSON:
{
    "_shards": {
        "failed": 0,
        "successful": 5,
        "total": 5
    },
    "hits": {
        "hits": [
            {
                "_id": "AVJOV5dSBy9sHt9l26MO",
                "_index": "ammonitrix_meta",
                "_score": 1.0,
                "_source": {
                    "name": "",
                    "tags": null
                },
                "_type": "register"
            }
        ],
        "max_score": 1.0,
        "total": 2
    },
    "timed_out": false,
    "took": 56
}
*/
// ElasticResponseMeta struct is the general struct of an elastic response
type ElasticResponseMeta struct {
	Took     float64               `json:"took"`
	TimedOut bool                  `json:"timed_out"`
	Shards   ElasticResponseShards `json:"_shards"`
	Hits     ElasticResponseHits   `json:"hits"`
}

func (o *ElasticResponseMeta) UnmarshalJSON(data []byte) error {
	var inter interface{}
	if err := json.Unmarshal(data, &inter); err != nil {
		return err
	}
	m := inter.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "took":
			o.Took = v.(float64)
		case "timedout":
			o.TimedOut = v.(bool)
		case "shards":
			o.Shards = v.(ElasticResponseShards)
		case "hits":
			o.Hits = v.(ElasticResponseHits)
		}
	}
	return nil
}

type ElasticResponseShards struct {
	Failed     float64 `json:"failed"`
	Successful float64 `json:"successful"`
	Total      float64 `json:"total"`
}

func (o *ElasticResponseShards) UnmarshalJSON(data []byte) error {
	var inter interface{}
	if err := json.Unmarshal(data, &inter); err != nil {
		return err
	}
	m := inter.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "failed":
			o.Failed = v.(float64)
		case "successful":
			o.Successful = v.(float64)
		case "total":
			o.Total = v.(float64)
		}
	}
	return nil
}

type ElasticResponseHits struct {
	MaxScore float64                 `json:"max_score"`
	Total    float64                 `json:"total"`
	Hits     ElasticResponseHitsData `json:"hits"`
}

func (o *ElasticResponseHits) UnmarshalJSON(data []byte) error {
	var inter interface{}
	if err := json.Unmarshal(data, &inter); err != nil {
		return err
	}
	m := inter.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "max_score":
			o.MaxScore = v.(float64)
		case "total":
			o.Total = v.(float64)
		case "hits":
			o.Hits = v.(ElasticResponseHitsData)
		}
	}
	return nil
}

type ElasticResponseHitsData struct {
	Index  string          `json:"_index"`
	Type   string          `json:"_type"`
	ID     string          `json:"_id"`
	Score  float64         `json:"_score"`
	Source ElasticMetadata `json:"_source"`
}

func (o *ElasticResponseHitsData) UnmarshalJSON(data []byte) error {
	var inter interface{}
	if err := json.Unmarshal(data, &inter); err != nil {
		return err
	}
	m := inter.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "_index":
			o.Index = v.(string)
		case "_type":
			o.Type = v.(string)
		case "_id":
			o.ID = v.(string)
		case "_score":
			o.Score = v.(float64)
		case "_source":
			o.Source = v.(ElasticMetadata)
		}
	}
	return nil
}
