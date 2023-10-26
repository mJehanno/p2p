package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	go func() {
		l, err := net.Listen("tcp", "127.0.0.1:9500")
		handleErr("fatal", "can't start tcp server", err)

		log.Info("Server listening on port 9500")
		for {
			c, err := l.Accept()
			handleErr("error", "can't accept connection", err)
			_, err = c.Write([]byte("hello you"))
			handleErr("error", "can't write to peer", err)
		}
	}()
	con, err := net.Dial("tcp", "tracker:8000")
	handleErr("fatal", "can't reach tracker node", err)
	var peerList string

	for len(peerList) <= 0 {
		_, err = con.Read([]byte(peerList))
		handleErr("fatal", "can't read peerlist from tracker", err)
		if len(peerList) > 0 {
			log.Info("received peer list")
			break
		}
	}

	con.Close()
	log.Infof("peer list received : %s", peerList)
	peers := strings.Split(peerList, "|")
	for _, p := range peers {
		co, err := net.Dial("tcp", p)
		handleErr("warn", "can't contact peer", err)
		_, err = co.Write([]byte("hey what's up ?"))
		handleErr("warn", "can't write to peer", err)

		var msg string
		_, err = co.Read([]byte(msg))
		handleErr("warn", "can't read from peer", err)
		fmt.Println(msg)
		defer co.Close()
	}
}

func handleErr(level, msg string, err error) {
	if err != nil {
		err = fmt.Errorf("%s : %w", msg, err)
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
