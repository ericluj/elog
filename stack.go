package log

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func stack(depth int) string {
	fmt.Println(111)
	pc, _, n, ok := runtime.Caller(1 + depth)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), n)
}
