package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/mxwuyin/constructorcheck"
)

func main() {
	singlechecker.Main(constructorcheck.Analyzer)
}
