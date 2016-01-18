package config

import "encoding/json"

type ElasticMetadata struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type ElasticData struct {
	Name      string      `json:"name"`
	CheckData interface{} `json:"check_data"`
}

type Datagram struct {
	Name      string      `json:"name"`
	Tags      []string    `json:"tags"`
	CheckData interface{} `json:"check_data"`
}

// ElasticResponse struct is the general struct of an elastic response
type ElasticResponse struct {
	Took     float64     `json:"took"`
	TimedOut bool        `json:"timed_out"`
	Shards   interface{} `json:"_shards"`
	Hits     interface{} `json:"hits"`
}

type ElasticResponseShards struct {
	Failed     float64 `json:"failed"`
	Successful float64 `json:"successful"`
	Total      float64 `json:"total"`
}

type ElasticResponseHits struct {
	MaxScore float64       `json:"max_score"`
	Total    float64       `json:"total"`
	Hits     []interface{} `json:"hits"`
}

type ElasticResponseHitsHits struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	ID     string      `json:"_id"`
	Score  float64     `json:"_score"`
	Source interface{} `json:"_source"`
}

// UnmarshalJSON for ElasticResponse struct
func (o *ElasticResponse) UnmarshalJSON(data []byte) error {
	var n interface{}
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	m := n.(map[string]interface{})
	for k, v := range m {
		if k == "took" {
			o.Took = v.(float64)
		} else if k == "timed_out" {
			o.TimedOut = v.(bool)
		} else if k == "_shards" {
			o.Shards = v.(map[string]interface{})
		} else if k == "hits" {
			o.Hits = v.(map[string]interface{})
		}
	}
	return nil
}

// UnmarshalJSON for ElasticResponseShards struct
func (o *ElasticResponseShards) UnmarshalJSON(data []byte) error {
	var n interface{}
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	m := n.(map[string]interface{})
	for k, v := range m {
		if k == "failed" {
			o.Failed = v.(float64)
		} else if k == "successful" {
			o.Successful = v.(float64)
		} else if k == "total" {
			o.Total = v.(float64)
		}
	}
	return nil
}

// UnmarshalJSON for ElasticResponseHits struct
func (o *ElasticResponseHits) UnmarshalJSON(data []byte) error {
	var n interface{}
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	m := n.(map[string]interface{})
	for k, v := range m {
		if k == "hits" {
			o.Hits = v.([]interface{})
		} else if k == "max_score" {
			o.MaxScore = v.(float64)
		} else if k == "total" {
			o.Total = v.(float64)
		}
	}
	return nil
}

// UnmarshalJSON for ElasticResponseHitsHits struct
func (o *ElasticResponseHitsHits) UnmarshalJSON(data []byte) error {
	var n interface{}
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	m := n.(map[string]interface{})
	for k, v := range m {
		if k == "_index" {
			o.Index = v.(string)
		} else if k == "_type" {
			o.Type = v.(string)
		} else if k == "_score" {
			o.Score = v.(float64)
		} else if k == "_id" {
			o.ID = v.(string)
		}
	}
	return nil
}
