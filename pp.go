package gocheckext

import (
	"strings"

	"github.com/go-test/deep"
)

var (
	// DeepLogErrors is a value for github.com/go-test/deep.LogErrors
	// while DeepEqualsPP.
	// In tests it's better to know when your DeepEqual check skip to
	// compare something.
	DeepLogErrors = true
	// DeepCompareUnexportedFields is a value for
	// github.com/go-test/deep.CompareUnexportedFields while DeepEqualsPP.
	// In tests it's usual to compare unexported fields.
	DeepCompareUnexportedFields = true
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
	defer func(deepLogErrors, deepCompareUnexportedFields bool) {
		deep.LogErrors = deepLogErrors
		deep.CompareUnexportedFields = deepCompareUnexportedFields
	}(deep.LogErrors, deep.CompareUnexportedFields)
	deep.LogErrors = DeepLogErrors
	deep.CompareUnexportedFields = DeepCompareUnexportedFields

	if diff := deep.Equal(params[0], params[1]); len(diff) > 0 {
		return false, "... ...\n  " + strings.Join(diff, "\n  ")
	}
	return true, ""
}
