package query

type ElasticServiceData struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string                          `json:"_index"`
			ID     string                          `json:"_id"`
			Score  float64                         `json:"_score"`
			Source ExampleMirroringTableForElastic `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type ExampleMirroringTableForElastic struct {
}

type ESBool struct {
	Must    []map[string]interface{} `json:"must,omitempty"`
	MustNot map[string]interface{}   `json:"must_not,omitempty"`
}

type ESQuery struct {
	Bool ESBool `json:"bool"`
}

type ESQueryWithAggsRequest struct {
	From  int                      `json:"from"`
	Size  int                      `json:"size"`
	Query ESQuery                  `json:"query"`
	Sort  []map[string]interface{} `json:"sort,omitempty"`
	Aggs  map[string]interface{}   `json:"aggs,omitempty"`
}

type ESBoolShould struct {
	Should []map[string]interface{} `json:"should"`
}
