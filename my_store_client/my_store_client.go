package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	cDefaultPort = "3000"
	cDefaultHost = "localhost"
	cFormat      = "$%d\r\n%s\r\n"
	cNumberInCMD = 3
	cHelpText    = `my_store_client GET NAME MANOJ
	SET NAME MYNAME
	DEL NAME`
)

var arguments []string

func main() {
	arguments = os.Args
	var hostAndPort string
	if len(arguments) == 4 {
		hostAndPort = fmt.Sprintf("%s:%s", cDefaultHost, cDefaultPort)
	} else {
		fmt.Println(cHelpText)
		return
	}

	fmt.Println("dialing, ", hostAndPort)
	c, err := net.Dial("tcp", hostAndPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(c, encodeToRedisPrrotocolSpec())
	message, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Print("->: " + message)
	return
}

func encodeToRedisPrrotocolSpec() string {
	formattedCommand := ""
	for i := 1; i <= cNumberInCMD; i++ {
		temp := arguments[i]
		formattedCommand += fmt.Sprintf(cFormat, len(temp), temp)
	}
	formattedCommand += fmt.Sprintf(cFormat, 4, "STOP")
	return formattedCommand
}
