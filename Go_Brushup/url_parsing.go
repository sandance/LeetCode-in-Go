package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	p, _ := u.User.Password()
	fmt.Println(p)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
