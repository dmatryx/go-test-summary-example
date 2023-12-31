package main

import (
	"github.com/dmatryx/go-test-summary-example/second-app/internal/alsountested"
	"github.com/dmatryx/go-test-summary-example/second-app/internal/secondexample"
)

func main() {
	print(secondexample.AddNumbers(alsountested.DoNothing(1), 2))
	print(secondexample.MirrorString("Testing"))
	print(secondexample.UntestedFunction("wooo", "bar"))
}
