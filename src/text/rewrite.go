package text

import (
	"fmt"
	"strings"

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

func Rewrite(text string, n int) []string {
	return getAnswers(rewrite, text, n)
}

func Longer(text string, n int) []string {
	return getAnswers(longer, text, n)
}

func Shorter(text string, n int) []string {
	return getAnswers(shorter, text, n)
}

func More(text string, parameter string, n int) []string {
	return getAnswersWithParameter(shorter, text, parameter, n)
}

func Less(text string, parameter string, n int) []string {
	return getAnswersWithParameter(shorter, text, parameter, n)
}

func getAnswers(prompt, text string, n int) []string {
	var out []string
	i := 0
	for i < n {
		candidate, err := oai.MakeCompletion(fmt.Sprintf(prompt, exceptions(out), text))
		if err != nil {
			panic(err)
		}

		if !validateCandidate(candidate) {
			continue
		}

		answer := extractAnswer(candidate)
		// we need to have exactly N completions
		if !validateAnswer(answer, out) {
			continue
		}

		out = append(out, answer)
		i++
	}

	return out
}

func getAnswersWithParameter(prompt, text, parameter string, n int) []string {
	var out []string
	i := 0
	for i < n {
		candidate, err := oai.MakeCompletion(fmt.Sprintf(prompt, parameter, exceptions(out), text))
		if err != nil {
			panic(err)
		}

		if !validateCandidate(candidate) {
			continue
		}

		answer := extractAnswer(candidate)
		// we need to have exactly N completions
		if !validateAnswer(answer, out) {
			continue
		}

		out = append(out, answer)
		i++
	}

	return out
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

func extractAnswer(candidate string) string {
	// contain <answer> and </answer>, one before another
	a := strings.Index(candidate, "<answer>") + len("<answer>")
	b := strings.Index(candidate, "</answer>")
	return candidate[a:b]
}

func validateCandidate(candidate string) bool {
	// contain <answer> and </answer>, one before another
	a := strings.Index(candidate, "<answer>")
	b := strings.Index(candidate, "</answer>")
	return a != -1 && a != b && a < b
}

func validateAnswer(answer string, exceptions []string) bool {
	// general answer validation
	// see if we already had this answer, or it is banned
	for _, e := range exceptions {
		if answer == e {
			return false
		}
	}
	return true
}
