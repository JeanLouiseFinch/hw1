package hw1

import (
	"fmt"
	"io"
	"time"

	"github.com/beevik/ntp"
)

const (
	ntpServRu = "0.ru.pool.ntp.org"
)

// Now - печатаем время с помощью стандартной библиотеки
func Now(out io.Writer) {
	fmt.Fprintf(out, "Системное время: %s\n", time.Now().Format(time.UnixDate))
}

//NowNTP - печатаем время по NTP серверу
func NowNTP(host string, out io.Writer) {
	if host == "" {
		host = ntpServRu
	}
	timeNTP, err := ntp.Time(host)
	if err != nil {
		fmt.Fprintf(out, "Что-то пошло не так(%v). Системное время: %s\n", err, time.Now().Format(time.UnixDate))
		return
	}
	fmt.Fprintf(out, "Время по NTP серверу: %s\n", timeNTP.Format(time.UnixDate))
}
