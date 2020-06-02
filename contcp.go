package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleTcpCon(c net.Conn) {
	for {
		reader, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Println("Avan poi pathu nimisham aagudhu")
			c.Close()
			return
		}

		message := string(reader)
		clientAddr := c.RemoteAddr().String()

		if len(clientAddr) == 0 {
			fmt.Println("Okay bye!!!")
			c.Close()
			os.Exit(1)
		}

		if strings.TrimSpace(message) == "STOP" {
			fmt.Println("seri kelambhu Mr." + clientAddr)
			c.Close()
			return
		}

		fmt.Println("-->", message+clientAddr)

		fmt.Fprintln(c, "received by server")
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("port number solluda ?!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("edho thappu nadakudhu inga")
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("edho thappu nadakudhu inga")
			return
		}

		go handleTcpCon(c)
	}

}
