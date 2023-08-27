package src

import (
	"fmt"
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
	same, err := text.Rewrite(request.Prompt, n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	longer, err := text.Longer(request.Prompt, n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	shorter, err := text.Shorter(request.Prompt, n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	more, err := text.More(request.Prompt, request.Parameter, n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	less, err := text.Less(request.Prompt, request.Parameter, n)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	fmt.Printf("Prompt: %s\n", request.Prompt)
	fmt.Printf("Same: %v\n", same)
	fmt.Printf("Longer: %v\n", longer)
	fmt.Printf("Shorter: %v\n", shorter)
	fmt.Printf("More %s: %v\n", request.Parameter, more)
	fmt.Printf("Less %s: %v\n", request.Parameter, less)

	// get differences between variants and text
	//diffs := make(map[int]map[int]string)
	//for i := range variants {
	//	diffs[i] = text.Difference(request.Prompt, variants)
	//}

	return c.JSON(http.StatusOK, nil)
}
