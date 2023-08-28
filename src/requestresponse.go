package src

// curl -X POST -H "Content-Type: application/json" -d '{"prompt":"I like to eat apples.", "n": 3}' http://localhost:8080/rewrite
// curl -X POST -H "Content-Type: application/json" -d '{"prompt":"I like to eat apples.", "parameter": "formal" ,"n": 1}' http://localhost:8080/rewrite

type Request struct {
	Parameter string `json:"parameter"`
	Prompt    string `json:"prompt"`
	N         int    `json:"n,omitempty"`
}

type Response struct {
	Prompt string `json:"prompt"`
	Parameter string `json:"parameter"`

	Same []string `json:"same,omitempty"`
	Longer []string `json:"longer,omitempty"`
	Shorter []string `json:"shorter,omitempty"`
	More []string `json:"more,omitempty"`
	Less []string `json:"less,omitempty"`
}
