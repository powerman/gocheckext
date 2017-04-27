package gocheckext

import (
	"testing"

	"gopkg.in/check.v1"
)

func init() {
	// TODO Can this be done using reflect to find and wrap all existing
	// checkers? But in this case it should be run by CountingTestingT
	// to make sure all checkers was imported/defined when this runs.
	// As a side-effect it may (harmlessly) double-wrap checkers if
	// CountingTestingT will be called multiple times for some reason.
	check.DeepEquals = CountChecker(check.DeepEquals)
	check.Equals = CountChecker(check.Equals)
	check.ErrorMatches = CountChecker(check.ErrorMatches)
	check.FitsTypeOf = CountChecker(check.FitsTypeOf)
	check.HasLen = CountChecker(check.HasLen)
	check.Implements = CountChecker(check.Implements)
	check.IsNil = CountChecker(check.IsNil)
	check.Matches = CountChecker(check.Matches)
	check.NotNil = CountChecker(check.NotNil)
	check.PanicMatches = CountChecker(check.PanicMatches)
	check.Panics = CountChecker(check.Panics)
}

var (
	checksPassed int
	checksFailed int
)

// CountingTestingT should be called instead of gocheck's TestingT.
// It will report count of passed/failed checks (like DeepEquals or IsNil)
// after usual TestingT report of passed/failed tests - so you can see how
// many real tests you've implemented so far.
//
//   func Test(t *testing.T) { gocheckext.CountingTestingT(t) }
//
// Then run `go test` as usually and you'll see count of checks after
// count of tests, for ex.:
//
//   OOPS: 20 passed, 3 skipped, 1 FAILED
//   Checks: 196 passed, 1 failed
//   --- FAIL: Test (0.56s)
//
//   OK: 21 passed, 3 skipped
//   Checks: 197 passed
//   --- PASS: Test (0.49s)
//
func CountingTestingT(t *testing.T) {
	check.TestingT(t)
	reportChecksCount()
}

func reportChecksCount() {
	if checksFailed == 0 {
		println("Checks:", checksPassed, "passed")
	} else {
		println("Checks:", checksPassed, "passed,", checksFailed, "failed")
	}
	checksPassed = 0
	checksFailed = 0
}

type countingChecker struct {
	info *check.CheckerInfo
	code func([]interface{}, []string) (bool, string)
}

func (c countingChecker) Info() *check.CheckerInfo { return c.info }
func (c countingChecker) Check(params []interface{}, names []string) (result bool, err string) {
	curPASS, curFAIL := checksPassed, checksFailed // protect against recursive counting checkers
	result, err = c.code(params, names)
	if result {
		checksPassed = curPASS + 1
		checksFailed = curFAIL
	} else {
		checksPassed = curPASS
		checksFailed = curFAIL + 1
	}
	return
}

// CountChecker wrap usual Checker to count it calls. You should use it
// to wrap all your custom (not gocheck's ones) checkers. Example:
//
//   import . "github.com/dropbox/godropbox/gocheck2"
//   // wrap custom checker
//   var myChecker = gocheckext.CountChecker(myCheckerType{…})
//   func init() {
//       // wrap checkers imported from godropbox
//       AlmostEqual = gocheckext.CountChecker(AlmostEqual)
//       BytesEquals = gocheckext.CountChecker(BytesEquals)
//       GreaterThan = gocheckext.CountChecker(GreaterThan)
//       ...
//   }
func CountChecker(f check.Checker) check.Checker {
	return countingChecker{
		info: f.Info(),
		code: f.Check,
	}
}

// NewCountingChecker is just a convenient shortcut for
// CountChecker(NewChecker(…)).
func NewCountingChecker(name string, params []string, code CheckFunc) check.Checker {
	return CountChecker(NewChecker(name, params, code))
}
