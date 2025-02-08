package infrastructure

type ResponseEmbedQuery struct {
	Vector []float64 `json:"vector"`
}

type RequestEmbedQuery struct {
	Text string `json:"text"`
}

type RequestScriptQuery struct {
	Text string `json:"text"`
}

type ResponseScriptQuery struct {
	Script ScriptBody `json:"script"`
}

type ScriptBody struct {
	Source []string  `json:"_source"`
	Query  QueryBody `json:"query"`
}

type QueryBody struct {
	ScriptScore ScriptScore `json:"script_score"`
}

type ScriptScore struct {
	Query  interface{} `json:"query"`
	Script Script      `json:"script"`
}

type Script struct {
	Source string       `json:"source"`
	Params ScriptParams `json:"params"`
}

type ScriptParams struct {
	QueryVector   []float64 `json:"query_vector"`
	TitleWeight   float64   `json:"title_weight"`
	AddressWeight float64   `json:"address_weight"`
}
