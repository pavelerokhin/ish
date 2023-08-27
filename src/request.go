package src

type RewriteRequest struct {
	Parameter string `json:"parameter"`
	Prompt    string `json:"prompt"`
	N         int    `json:"n,omitempty"`
}

// curl -X POST -H "Content-Type: application/json" -d '{"prompt":"I like to eat apples.", "n": 3}' http://localhost:8080/rewrite
// curl -X POST -H "Content-Type: application/json" -d '{"prompt":"I like to eat apples.", "parameter": "formal" ,"n": 1}' http://localhost:8080/rewrite
