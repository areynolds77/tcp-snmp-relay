package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	g "github.com/soniah/gosnmp"
)

func sendTrap(temp string) {
	fmt.Printf("Received Data: %s\n", temp)
	s := strings.Split(temp, ",")
	TargetIP := s[0]
	TargetOID := s[1]
	TargetCommunity := s[2]
	//TargetValue := s[2]
	fmt.Println(TargetIP)
	fmt.Println(TargetOID)
	//fmt.Println(TargetValue)

	g.Default.Target = TargetIP
	g.Default.Port = 162
	g.Default.Version = g.Version2c
	g.Default.Community = TargetCommunity
	g.Default.Logger = log.New(os.Stdout, "", 0)

	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	pdu := g.SnmpPDU{
		Name:  TargetOID,
		Type:  g.ObjectIdentifier,
		Value: TargetOID,
	}

	trap := g.SnmpTrap{
		Variables: []g.SnmpPDU{pdu},
	}

	_, err = g.Default.SendTrap(trap)
	if err != nil {
		log.Fatalf("SendTrap() err: %v", err)
	}

}
func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		go sendTrap(temp)
	}
	c.Close()
}

func main() {
	PORT := ":8162"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
