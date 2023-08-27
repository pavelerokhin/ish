package src

type RewriteRequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n,omitempty"`
}
