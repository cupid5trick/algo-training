package main

import (
	"os"
	"strings"
	"testing"
	. "github.com/cupid5trick/algotrain/learn/testutil"
)

func TestC(t *testing.T) {
	checkPoints := [][2]string{
		{}
	}
	AssertEqualRsults(t, checkPoints, run)
}