// Package gocheckext provides extensions for github.com/go-check/check.
package gocheckext

import "gopkg.in/check.v1"

// CheckFunc is a gocheck's Checker.Check function.
type CheckFunc func([]interface{}, []string) (bool, string)

type aChecker struct {
	Desc *check.CheckerInfo
	Func CheckFunc
}

// Info implements gocheck's Checker.Info.
func (c aChecker) Info() *check.CheckerInfo {
	return c.Desc
}

// Check implements gocheck's Checker.Check.
func (c aChecker) Check(params []interface{}, names []string) (result bool, err string) {
	return c.Func(params, names)
}

// NewChecker makes it easier to define your own checkers.
func NewChecker(name string, params []string, code CheckFunc) check.Checker {
	return aChecker{
		&check.CheckerInfo{Name: name, Params: params},
		code,
	}
}
