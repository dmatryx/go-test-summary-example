package main

import (
	"github.com/dmatryx/go-test-summary-example/internal/example"
	"github.com/dmatryx/go-test-summary-example/internal/untested"
)

func main() {
	print(example.AddNumbers(untested.DoNothing(1), 2))
	print(example.MirrorString("Testing"))
	print(example.UntestedFunction("wooo", "bar"))
}
