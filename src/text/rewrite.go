package text

import (
	"fmt"

	"ish/src/oai"
)

const (
	rewrite = "Rewrite the following sentence having the same length " +
		"your answer must be different from sentences inside <no> XML tags " +
		"<no>%s</no>" +
		"%s\n"
	longer = "Rewrite the following sentence a bit longer " +
		"your answer must be different from sentences inside <no> XML tags " +
		"<no>%s</no>" +
		"%s\n"
	shorter = "Rewrite the following sentence a bit shorter " +
		"your answer must be different from sentences inside <no> XML tags" +
		"<no>%s</no>" +
		"%s\n"
	more = "Rephrase this sentence to make it clearer and more %s" +
		"your answer must be different from sentences inside <no> XML tags" +
		"<no>%s</no>" +
		"%s\n"
	less = "Rephrase this sentence to make it clearer and less %s" +
		"your answer must be different from sentences inside <no> XML tags" +
		"<no>%s</no>" +
		"%s\n"
)

func Rewrite(text string, n int) ([]string, error) {
	var out []string

	for i := 0; i < n; i++ {
		variant, err := oai.MakeCompletion(fmt.Sprintf(rewrite, out, text))
		// we need to have exactly N completions
		if err != nil {
			panic(err)
		}
		out = append(out, variant)
	}

	return out, nil
}

func Longer(text string, n int) ([]string, error) {
	var out []string

	for i := 0; i < n; i++ {
		variant, err := oai.MakeCompletion(fmt.Sprintf(longer, out, text))
		// we need to have exactly N completions
		if err != nil {
			panic(err)
		}
		out = append(out, variant)
	}

	return out, nil
}

func Shorter(text string, n int) ([]string, error) {
	var out []string

	for i := 0; i < n; i++ {
		variant, err := oai.MakeCompletion(fmt.Sprintf(shorter, out, text))
		// we need to have exactly N completions
		if err != nil {
			panic(err)
		}
		out = append(out, variant)
	}

	return out, nil
}

func More(text string, parameter string, n int) ([]string, error) {
	var out []string

	for i := 0; i < n; i++ {
		x := fmt.Sprintf(more, parameter, out, text)
		variant, err := oai.MakeCompletion(x)
		// we need to have exactly N completions
		if err != nil {
			panic(err)
		}
		out = append(out, variant)
	}

	return out, nil
}

func Less(text string, parameter string, n int) ([]string, error) {
	var out []string

	for i := 0; i < n; i++ {
		variant, err := oai.MakeCompletion(fmt.Sprintf(less, parameter, out, text))
		// we need to have exactly N completions
		if err != nil {
			panic(err)
		}
		out = append(out, variant)
	}

	return out, nil
}
