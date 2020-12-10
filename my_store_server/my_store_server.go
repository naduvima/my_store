package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

const (
	cDefaultPort = "3000"
	cFormat      = "$%d\r\n%s\r\n"
)

type cmdInfo struct {
	CmdLen int    `json:"length_of_command"`
	Cmd    string `json:"cmd"`
	KeyLen int    `json:"length_of_key"`
	Key    string `json:"key"`
	ValLen int    `json:"length_of_value"`
	Val    string `json:"value"`
}

var supportedCommands = []string{"GET", "DEL", "SET", "STOP"}

func main() {
	arguments := os.Args

	var PORT string
	if len(arguments) == 1 {
		PORT = ":" + cDefaultPort
	} else {
		PORT = ":" + arguments[1]
	}

	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	fmt.Println("MY Store ready and listening at: ", cDefaultPort)
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println("Reading from remote address: ", c.RemoteAddr())
	temp := make([]byte, 1024)
	for {
		_, err := c.Read(temp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(temp))
		break
	}
	/*here is the funniest part I am constructing a json
	and sending it across
	*/
	responseJSON, _ := json.Marshal(deCodetoMap(temp))
	fmt.Println(string(responseJSON))
	c.Write([]byte("Recieved the following for client's record\n"))
	c.Write(responseJSON)
	c.Write([]byte(" Closing Client..\n"))
	c.Close()
	fmt.Println("Closed.", c.RemoteAddr())
}

func deCodetoMap(temp []byte) cmdInfo {
	cInfo := cmdInfo{}
	str := string(temp)
	fmt.Printf(str)
	fmt.Sscanf(
		str,
		cFormat+cFormat+cFormat,
		&cInfo.CmdLen, &cInfo.Cmd,
		&cInfo.KeyLen, &cInfo.Key,
		&cInfo.ValLen, &cInfo.Val,
	)
	fmt.Println(cInfo)
	return cInfo
}
