package main

import (
	"os"
	"time"

	"github.com/nslee333/go_with_tdd/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
