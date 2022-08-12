package main

import (
	"net"
	"fmt"
	"io/ioutil"
	"log"
)

func handleError (err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func makeRequest (path string) string {
	resolvedIp, err := net.ResolveTCPAddr("tcp", "localhost:8080")

	connection, err := net.DialTCP("tcp", nil, resolvedIp)

	handleError(err)

	message := "GET " + path + " HTTP/1.0\r\nHost: jamesg.blog\r\n\r\n"

	_, err = connection.Write([]byte(message))

	handleError(err)

	result, err := ioutil.ReadAll(connection)

	handleError(err)

	return string(result)
}

func main () {
	fmt.Println(makeRequest("/"))
}