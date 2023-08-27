package src

import (
	"ish/src/text"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Rewrite(c echo.Context) error {
	var request RewriteRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if request.Prompt == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Prompt is required")
	}
	n := max(request.N, 1)

	// get variants of text
	var variants []string

	for i := 0; i < n; i++ {
		variants = append(variants, text.Rewrite(request.Prompt))
	}

	// get differences between variants and text
	diffs := make(map[int]map[int]string)
	for i := range variants {
		diffs[i] = text.Difference(request.Prompt, variants)
	}

	return nil
}
