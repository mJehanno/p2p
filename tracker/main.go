package main

import (
	"fmt"
	"net"

	"github.com/charmbracelet/log"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	handleErr("fatal", "can't start tcp server", err)
	peers := []net.Conn{}

	peerlist := ""
	for {
		con, err := l.Accept()
		handleErr("error", "can't accept connection", err)

		peers = append(peers, con)
		for _, c := range peers {
			peerlist += c.RemoteAddr().String() + "|"
		}
		log.Infof("current value of peerlist : %s", peerlist)
		_, err = con.Write([]byte(peerlist))
		handleErr("error", "can't write to peer", err)
	}
}

func handleErr(level, msg string, err error) {
	if err != nil {
		err := fmt.Errorf("%s : %w", msg, err)
		switch level {
		case "fatal":
			log.Fatal(err)
		case "error":
			log.Error(err)
		case "warn":
			log.Warn(err)
		}
	}
}
