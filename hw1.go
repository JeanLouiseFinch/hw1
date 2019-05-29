package hw1

import (
	"fmt"
	"io"
	"time"
)

// Now - печатаем время с помощью стандартной библиотеки
func Now(out io.Writer) {
	fmt.Fprintf(out, "%s\n", time.Now().Format(time.UnixDate))
}
