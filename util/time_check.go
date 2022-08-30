package util

import (
	"fmt"
	"strings"
	"time"
)

func TimeCheck(fn func(), args ...string) {
	start := time.Now()
	fn()
	fmt.Println(time.Since(start).Milliseconds(), "ms", strings.Join(args, " "))
}
