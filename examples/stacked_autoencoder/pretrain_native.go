//go:build native
// +build native

package main

import (
	. "github.com/jeremypansier/gorgonia"
	"gonum.org/v1/gonum/blas/gonum"
)

func init() {
	Use(gonum.Implementation{})
}
