package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/ssh"
)

type ClientConf struct{
	Username string
	Password string
	SshServer string
	SshPort string
	SshHttpProxyServer string
	SshHttpProxyServerPort string
}

func SshHttpProxy(c *ClientConf) {
	var hostKey ssh.PublicKey
	config := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	// Dial your ssh server.
	conn, err := ssh.Dial("tcp", c.SshServer + ":" + c.SshPort, config)
	if err != nil {
		log.Fatal("unable to connect: ", err)
	}
	defer conn.Close()

	// Request the remote side to open port 8080 on all interfaces.
	l, err := conn.Listen("tcp",c.SshHttpProxyServer + ":" + c.SshHttpProxyServerPort)
	if err != nil {
		log.Fatal("unable to register tcp forward: ", err)
	}
	defer l.Close()

	// Serve HTTP with your SSH server acting as a reverse proxy.
	http.Serve(l, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Hello world!\n")
	}))
}

func main(){
	config := ClientConf{
		Username: "sam",
		Password: "3l0r@cle!.",
		SshServer: "127.0.0.1",
		SshPort: "2222",
		SshHttpProxyServer: "0.0.0.0",
		SshHttpProxyServerPort: "8080",
	}
	SshHttpProxy(&config)
}
