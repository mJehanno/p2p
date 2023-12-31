package main

import (
	"fmt"
	"net"
	"regexp"

	"github.com/charmbracelet/log"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	handleErr("fatal", "can't start tcp server", err)
	peers := map[string]net.Conn{}

	for {
		con, err := l.Accept()
		handleErr("error", "can't accept connection", err)

		go handleConnection(con, peers)
	}
}

func handleConnection(con net.Conn, peers map[string]net.Conn) {
	peers[con.RemoteAddr().String()] = con
	peerlist := ""
	_, err := con.Write([]byte(con.RemoteAddr().String()))
	handleErr("error", "can't write to peer", err)
	for _, c := range peers {
		reg := regexp.MustCompile("/:*$/")
		addr := reg.ReplaceAllString(c.RemoteAddr().String(), ":9500")
		peerlist += addr + "|"
	}
	log.Infof("current value of peerlist : %s", peerlist)
	_, err = con.Write([]byte(peerlist))
	handleErr("error", "can't write to peer", err)

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
