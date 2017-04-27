package gocheckext

import (
	"strings"

	"github.com/go-test/deep"
)

// DeepEqualsPP works like gocheck's DeepEquals but also pretty-print a
// list of differences.
//
// You can easily make all your current tests use DeepEqualsPP with:
//
//   func init() {
//       DeepEquals = gocheckext.DeepEqualsPP
//   }
var DeepEqualsPP = NewCountingChecker(
	"DeepEqualsPP", []string{"obtained", "expected"}, deepEquals,
)

func deepEquals(params []interface{}, names []string) (result bool, err string) {
	if diff := deep.Equal(params[0], params[1]); len(diff) > 0 {
		return false, "... ...\n  " + strings.Join(diff, "\n  ")
	}
	return true, ""
}
