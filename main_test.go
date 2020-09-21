package main

import (
	"net"
	"testing"
)

func Test_Router(t *testing.T) {
	err := router("GET", "/")

	if err != nil {
		t.Fatal("Expected nil error for GET request")
	}

	err = router("POST", "/")

	if err == nil {
		t.Fatalf("Expected nil, Got %s", err.Error())
	}
}

func Test_Handler(t *testing.T) {
	server, client := net.Pipe()

	go handleConnection(server)

	client.Write([]byte("GET /\n"))
	resp, _ := read(client)

	expected := "You asked for /\n"

	if resp != expected {
		t.Fatalf("Expected (%s), Got (%s)", expected, resp)
	}

	client.Close()
}

func Test_Prase(t *testing.T) {
	method, url := parseHttpString("GET / ")

	if method != "GET" || url != "/" {
		t.Fatalf("Expected (%s, %s), Got (%s, %s)", "GET", "/", method, url)
	}
}
