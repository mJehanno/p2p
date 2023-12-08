package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	// Launching server on peer
	go func() {
		l, err := net.Listen("tcp", "0.0.0.0:"+os.Getenv("PORT"))
		handleErr("fatal", "can't start tcp server", err)
		log.Infof("tcp server listening on port :%s", os.Getenv("PORT"))

		for {
			c, err := l.Accept()
			handleErr("error", "can't accept connection", err)

			handleConnection(c)

		}
	}()

	// Reaching tracker
	con, err := net.Dial("tcp", "tracker:8000")
	handleErr("fatal", "can't reach tracker node", err)

	var stringedPeerList string
	msg := make([]byte, 1024)
	_, err = con.Read(msg)
	currentIp := string(msg)
	log.Infof("current ip : %s", currentIp)
	peerLen := 0
	peers := []string{}

	for {
		_, err = con.Read(msg)
		handleErr("fatal", "can't read peerlist from tracker", err)

		stringedPeerList = string(msg)

		if len(stringedPeerList) != peerLen {
			peerList := strings.Split(stringedPeerList, "|")

			peerList = slices.DeleteFunc(peerList, func(ip string) bool {
				return strings.Contains(ip, currentIp) || strings.Contains(currentIp, ip)
			})
			peerLen = len(stringedPeerList)
			peers = peerList
		}

		if len(peers) > 0 {
			log.Infof("peer list received : %s", peers)
			break
		}
	}

	con.Close()

	for _, p := range peers {
		reg := regexp.MustCompile(`:[0-9]*`)
		res := reg.ReplaceAll([]byte(p), []byte(""))
		co, err := net.Dial("tcp", fmt.Sprintf("%s:%s", string(res), os.Getenv("PORT")))
		handleErr("fatal", "can't contact peer", err)
		if co != nil {
			msg := make([]byte, 1024)
			_, err = co.Read([]byte(msg))
			handleErr("warn", "can't read from peer", err)
			log.Infof("received msg : %s", msg)

			_, err = co.Write([]byte("hey what's up ?"))
			handleErr("warn", "can't write to peer", err)

		}
		defer co.Close()
	}
}

func handleConnection(co net.Conn) {
	_, err := co.Write([]byte("hello you"))
	handleErr("error", "can't write to peer", err)

	msg := make([]byte, 1024)

	_, err = co.Read(msg)
	handleErr("error", "can't read from peer", err)

	if err == nil {
		log.Infof("received msg : %s", msg)
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
