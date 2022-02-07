//go:build !native
// +build !native

package main

import (
	. "github.com/jeremypansier/gorgonia"
	"github.com/jeremypansier/gorgonia/blase"
)

func init() {
	Use(blase.Implementation())
}
