package main

import (
	"github.com/keithbro-imx/go-func-doc-prefix/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
