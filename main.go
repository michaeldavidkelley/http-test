package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

func router(method, url string) error {
	switch method {
	case "GET":
		return nil
	}

	return errors.New("Could not handle request")
}

func read(c net.Conn) (string, error) {
	return bufio.NewReader(c).ReadString('\n')
}

func write(c net.Conn, str string) {
	c.Write([]byte(str))
}

func parseHttpString(content string) (string, string) {
	parts := strings.Split(content, " ")
	return parts[0], parts[1]
}

func handleConnection(c net.Conn) {
	defer c.Close()

	str, err := read(c)
	if err != nil {
		c.Close()
	}

	method, url := parseHttpString(str)

	if err := router(method, url); err != nil {
		return
	}

	write(c, fmt.Sprintf("You asked for %s", url))
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Could not listen")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Could not accept connection")
			continue
		}
		go handleConnection(conn)
	}
}
