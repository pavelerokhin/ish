package src

import (
	"fmt"
	"ish/src/text"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

func Rewrite(c echo.Context) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if request.Prompt == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Prompt is required")
	}
	n := max(request.N, 1)

	// get variants of text
	var same, longer, shorter, more, less []string
	wg := &sync.WaitGroup{}
	wg.Add(5)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		same = text.Rewrite(request.Prompt, n)
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		longer = text.Longer(request.Prompt, n)
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		shorter = text.Shorter(request.Prompt, n)
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		more = text.More(request.Prompt, request.Parameter, n)
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		less = text.Less(request.Prompt, request.Parameter, n)
	}(wg)

	wg.Wait()
	fmt.Printf("Prompt: %s\n", request.Prompt)
	fmt.Printf("Same: %v\n", same)
	fmt.Printf("Longer: %v\n", longer)
	fmt.Printf("Shorter: %v\n", shorter)
	fmt.Printf("More %s: %v\n", request.Parameter, more)
	fmt.Printf("Less %s: %v\n", request.Parameter, less)

	resp := Response{
		Prompt:    request.Prompt,
		Parameter: request.Parameter,

		Same:    same,
		Longer:  longer,
		Shorter: shorter,
		More:    more,
		Less:    less,
	}
	// encode json resp e send to client
	return c.JSON(http.StatusOK, resp)
}
