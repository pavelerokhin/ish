package text

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func Rewrite(text string) string {
	var out string

	return out
}

func addRewritePrompt(request *openai.CompletionRequest) {
	request.Prompt = fmt.Sprintf("Rewrite the following sentence in a clear anf friendly way:%s\n", request.Prompt)
}
