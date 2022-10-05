package months

import (
	"errors"
	"time"

	"github.com/yudgxe/mini-sort/cstrings"
)

var layouts = []string{
	"January",
	"Jan",
	"1",
	"01",
}

func Parse(in string) (int, error) {
	for _, v := range layouts {
		t, err := time.Parse(v, in)
		if err == nil {
			return int(t.Month()), nil
		}
	}

	return -1, errors.New("Cannot parse \"" + in + "\" as " + cstrings.BuildFromWords(", ", layouts...))
}
