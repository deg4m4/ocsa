package core

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
)

func (o *Ocsa) getListner() net.Listener {
	if o.tls {
		cert, err := tls.LoadX509KeyPair(o.tlsOptions.cert, o.tlsOptions.key)
		if err != nil {
			fmt.Printf("Error loading cert: %s\n", err)
		}

		config := tls.Config{Certificates: []tls.Certificate{cert}}
		config.Rand = rand.Reader

		listener, err := tls.Listen("tcp", fmt.Sprintf("%s:%d", o.host, o.port), &config)
		if err != nil {
			fmt.Printf("Error listening: %s\n", err)
		}

		if o.verbose {
			fmt.Printf("Server listening on %s:%d\n tls", o.host, o.port)
		}

		return listener

	} else {

		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", o.host, o.port))
		if err != nil {
			fmt.Printf("Error listening: %s\n", err)
		}

		if o.verbose {
			fmt.Printf("Server listening on %s:%d\n", o.host, o.port)
		}

		return listener

	}
}
