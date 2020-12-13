package main

import (
	"fmt"
	"my_store/redis"
	"net"
	"os"
)

const (
	cDefaultPort = "3000"
	cDefaultHost = "localhost"
	cHelpText    = `my_store_client GET NAME MANOJ
	SET NAME MYNAME
	DEL NAME`
)

var arguments []string
var argNumbers map[string]int

func main() {
	arguments = os.Args
	var hostAndPort string
	argNumbers = map[string]int{"GET": 2, "SET": 3, "DEL": 2}
	hostAndPort = fmt.Sprintf("%s:%s", cDefaultHost, cDefaultPort)
	if len(arguments) == 1 || len(arguments[1:]) != argNumbers[arguments[1]] {
		fmt.Println(cHelpText)
		return
	}

	fmt.Println("dialing, ", hostAndPort)
	c, err := net.Dial("tcp", hostAndPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(c, encodeToRedisProtocolSpec())
	message := make([]byte, 1024)
	_, err = c.Read(message)
	msg, err := redis.RESPhandler(message)
	if err == nil {
		fmt.Print("->: " + string(msg))
	} else {
		fmt.Print("->: " + string(err.Error()))
	}
	return
}

func encodeToRedisProtocolSpec() string {
	return redis.EncodeWordsToRedisSpec(arguments[1:], argNumbers[arguments[1]])
}
