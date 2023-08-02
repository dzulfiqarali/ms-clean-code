package main

//go:generate go run github.com/google/wire/cmd/wire

func main() {
	//test
	server := InitializeService()
	server.Serve()
}
