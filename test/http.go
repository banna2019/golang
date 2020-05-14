package main

import "net/http"

func main() {
	http.listenAndServer("127.0.0.1",handler);
}