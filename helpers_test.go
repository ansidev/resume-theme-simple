package theme

import (
	"testing"
)

func TestFuncs(t *testing.T) {
	funcs := Funcs()

	funcNames := []string{
		"bootstrapIcon",
		"joinArray",
		"markdown",
		"displayUrl",
		"toLowerCase",
		"year",
	}

	for _, funcName := range funcNames {
		if _, ok := funcs[funcName]; !ok {
			t.Errorf("Theme function %s was not exists", funcName)
		}
	}
}
