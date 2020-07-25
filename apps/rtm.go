// https://golangbot.com/learn-golang-series/
// https://github.com/coocos/catastrophe
// https://github.com/mmcgrana/gobyexample/tree/master/examples
// https://golang.org/doc/effective_go.html

package main

import (
	"log"

	"rimortm/udpserver"
	"rimortm/webserver"
)

func main() {
	log.Println("Starting RTM")

	_ = udpserver.New("127.0.0.1", 6000)
	_ = webserver.New("127.0.0.1", 8080)

	// Block forever while goroutines run as needed
	select {}
}
