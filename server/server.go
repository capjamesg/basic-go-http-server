package main

import (
	"net"
	"log"
	"fmt"
	"strings"
	"time"
	"os"
)

func handleError (err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLinkHeaders () string {
	LinkHeaders := []string{
		"link: <http://localhost:8080/>; rel=\"self\"",
		"link: <http://localhost:8080/>; rel=\"micropub\"",
	}
	return strings.Join(LinkHeaders, "\r\n")
}

const (
	HTTP_VERSION = "HTTP/1.1 "
	Success = "200 OK\r\n"
	NotFound = "404 Not Found\r\n"
	StandardHeaders = "Content-Type: text/html\r\nX-Powered-by: Coffee and code\r\nX-Favourite-Coffee-Brewer: Aeropress\r\n"
)

func packageResponse (content string, status string) string {
	contentLength := "Content-Length: " + fmt.Sprintf("%d", len([]byte(content)))
	date := fmt.Sprintf("Date: %s", time.Now().Format(time.RFC1123))
	
	return (
		HTTP_VERSION +
		status +
		date + "\r\n" +
		StandardHeaders +
		contentLength + "\r\n" +
		getLinkHeaders() + "\r\n\r\n" +
		content + "\r\n")
}

func handleConnection (connection net.Conn) {
	data := make([]byte, 1024)

	contents, err := connection.Read(data)

	handleError(err)

	request := string(data[:contents])

	// get file path
	resource := strings.Split(request, " ")[1]

	fmt.Println("Connection established with", connection.RemoteAddr())

	allowedResources := map[string]string{
		"/": "/index.html",
	}

	var response string
	var status string

	if _, ok := allowedResources[resource]; ok {
		fileContents, err := os.ReadFile("pages" + allowedResources[resource])
		handleError(err)

		if resource == "/" {
			response = string(fileContents)
			status = Success
		}
	} else {
		response = "404!"
		status = NotFound
	}

	sendBack := packageResponse(response, status)

	connection.Write([]byte(sendBack))

	connection.Close()
}

func main () {
	resolvedIp, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	handleError(err)

	listen, err := net.ListenTCP("tcp", resolvedIp)
	handleError(err)

	for {
		connection, err := listen.Accept()
	
		handleError(err)

		go handleConnection(connection)
	}
}

// SELECT path FROM logs WHERE date <= today AND date >= (today - 90 days) GROUP BY path ORDER BY COUNT(path) DESC LIMIT 10;