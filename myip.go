package main

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

func main() {
	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range ifaces {
		if i.Name == "lo" {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Printf("%v: %v\n", i.Name, ip)
			// process IP address
		}
	}
	fmt.Println("External: " + external())
}

//dig +short myip.opendns.com @resolver1.opendns.com
func external() (ext string) {
	ext = "Not found"
	target := "myip.opendns.com"
	server := "resolver1.opendns.com"

	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(target+".", dns.TypeA)
	r, _, err := c.Exchange(&m, server+":53")
	if err != nil {
		//log.Println(err)
		return
	}
	for _, ans := range r.Answer {
		Arecord := ans.(*dns.A)
		ext = Arecord.A.String()
	}
	return
}
