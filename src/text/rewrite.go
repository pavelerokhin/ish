package text

import (
	"fmt"

	"ish/src/oai"
)

const (
	exception      = "<no>%s</no>\n"
	exceptionFrame = "<answer> has to be different from phrases inside <no> XML tags.\n"
	rewrite        = `Rewrite a sentence inside <sentence> XML tags of same length.
Place the answer inside <answer> XML.
%s
<sentence>%s</sentence>\n`
	longer = `Rewrite a sentence inside <sentence> XML tags as a longer sentence.
Place the answer inside <answer> XML.
%s
<sentence>%s</sentence>\n`
	shorter = `Rewrite a sentence inside <sentence> XML tags as a shorter sentence.
Place the answer inside <answer> XML. 
%s
<sentence>%s</sentence>\n`
	more = `Rewrite a sentence inside <sentence> XML tags. It must be more %s.
Place the answer inside <answer> XML. 
%s
<sentence>%s</sentence>\n`
	less = `Rewrite a sentence inside <sentence> XML tags. It must be less %s.
Place the answer inside <answer> XML. 
%s
<sentence>%s</sentence>\n`
)

func Rewrite(text string, n int) ([]string, error) {
	var out []string

	for i := 0; i < n; i++ {
		variant, err := oai.MakeCompletion(fmt.Sprintf(rewrite, exceptions(out), text))
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
		variant, err := oai.MakeCompletion(fmt.Sprintf(longer, exceptions(out), text))
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
		variant, err := oai.MakeCompletion(fmt.Sprintf(shorter, exceptions(out), text))
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
		x := fmt.Sprintf(more, parameter, exceptions(out), text)
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
		variant, err := oai.MakeCompletion(fmt.Sprintf(less, parameter, exceptions(out), text))
		// we need to have exactly N completions
		if err != nil {
			panic(err)
		}
		out = append(out, variant)
	}

	return out, nil
}

func exceptions(ee []string) string {
	out := ""
	if len(ee) == 0 {
		return out
	}
	out = exceptionFrame
	for _, e := range ee {
		out += fmt.Sprintf(exception, e)
	}

	return out
}
