package api

type payload struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type document struct {
	Doc   string `json:"-"`
	Title string `json:"title"`
	DocID string `json:"DocID"`
}

type token struct {
	Line string `json:"-"`
}
