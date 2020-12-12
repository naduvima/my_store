package main

import (
	"fmt"
	"my_store/redis"
	"net"
	"os"
)

const (
	cDefaultPort = "3000"
)

type cmdInfo struct {
	CmdLen int    `json:"length_of_command"`
	Cmd    string `json:"cmd"`
	KeyLen int    `json:"length_of_key"`
	Key    string `json:"key"`
	ValLen int    `json:"length_of_value"`
	Val    string `json:"value"`
}

//GloablStore holds the store
var GloablStore = make(map[string]string)

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
		//fmt.Println(string(temp))
		break
	}
	/*here is the funniest part I am constructing a json
	and sending it across
	*/

	Decoded := deCodetoMap(temp)
	//responseJSON, _ := json.Marshal(Decoded)
	//fmt.Println(string(responseJSON))
	//c.Write([]byte("Recieved the following for client's record: "))
	//c.Write(responseJSON)

	switch Decoded.Cmd {
	case "SET":
		go buildGlobalStore(Decoded.Key, Decoded.Val)
	case "GET":
		val := []string{getFromGlobalStore(Decoded.Key)}
		encoded := redis.EncodeWordsToRedisSpec(val, 1)
		c.Write([]byte(encoded))
	case "DEL":
		go deleteFromGlabalStore(Decoded.Key)
	default:
		val := []string{"Unknown format"}
		encoded := redis.EncodeWordsToRedisSpec(val, 1)
		c.Write([]byte(encoded))
	}

	c.Close()
	fmt.Println("Closed.", c.RemoteAddr())
}

func deCodetoMap(temp []byte) cmdInfo {
	cInfo := cmdInfo{}
	str := string(temp)
	fmt.Sscanf(
		str,
		redis.CRedisFormat+redis.CRedisFormat+redis.CRedisFormat,
		&cInfo.CmdLen, &cInfo.Cmd,
		&cInfo.KeyLen, &cInfo.Key,
		&cInfo.ValLen, &cInfo.Val,
	)
	fmt.Println(cInfo)
	return cInfo
}

func buildGlobalStore(key string, value string) {
	GloablStore[key] = value
}

func getFromGlobalStore(key string) string {
	return GloablStore[key]
}

func deleteFromGlabalStore(key string) {
	delete(GloablStore, key)
}
