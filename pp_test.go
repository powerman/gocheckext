package gocheckext

import . "gopkg.in/check.v1"

type PPSuite struct{}

var _ = Suite(&PPSuite{})

// Work around https://github.com/go-check/check/pull/97
func not(c Checker) Checker {
	return aChecker{
		c.Info(),
		func(params []interface{}, names []string) (bool, string) {
			result, _ := c.Check(params, names)
			return !result, ""
		},
	}
}

func (s *PPSuite) TestDeepEqualsPP(c *C) {
	deepEqualsPP := CountChecker(CountChecker(NewCountingChecker(
		"DeepEqualsPP", []string{"obtained", "expected"}, deepEquals,
	)))
	c.Check(DeepEqualsPP, DeepEqualsPP, deepEqualsPP)
	c.Check(DeepEqualsPP, not(DeepEqualsPP), DeepEquals)
	c.Check(deepEqualsPP, deepEqualsPP, deepEqualsPP)
}
