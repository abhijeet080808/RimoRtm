// https://golangbot.com/learn-golang-series/
// https://github.com/coocos/catastrophe
// https://github.com/mmcgrana/gobyexample/tree/master/examples
// https://golang.org/doc/effective_go.html

package main

import (
	"rimortm/webserver"
)

func main() {
	s := webserver.New("127.0.0.1", 8080)
	s.Start()
}
