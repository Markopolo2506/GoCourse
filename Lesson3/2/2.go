package main

import "fmt"

type Configuration struct {
	Val   string
	Proxy struct {
		Address string
		Port    string
	}
}

func main() {

	c := &Configuration{
		Val: "test",
	}
	c.Proxy.Address = `127.0.0.1`
	c.Proxy.Port = `8080`

	fmt.Printf("%+v", c)
}
