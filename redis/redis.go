package redis

import "fmt"

/*CRedisFormat holds the redis communication format
refer https://redis.io/topics/protocol*/
const (
	CRedisFormat = "$%d\r\n%s\r\n"
	CRedisSucces = "+OK\r\n"
	CRedisErrors = "-Error %s\r\n"
	CRedis       = "$-1\r\n" //Null Bulk String, send when key not found
	CREdisTrue   = ":1\r\n"  // Return this when operation unsuccessful and no descrptive error message to send a CRedisErrors
)

/*EncodeWordsToRedisSpec pass array of words, they will be convered to a single string of redis spec format
Example:-
GET NAME HELLO will be converted to:-
$3\r\nGET\r\n$4\r\nNAME\r\n$5\r\nHELLO\r\n
*/
func EncodeWordsToRedisSpec(words []string, countOfWords int) string {
	formattedCommand := ""
	for i := 0; i < countOfWords; i++ {
		temp := words[i]
		formattedCommand += fmt.Sprintf(CRedisFormat, len(temp), temp)
	}
	return formattedCommand
}

/*EncodeErrorToRedisSpec formats to error spec*/
func EncodeErrorToRedisSpec(errText string) string {
	return fmt.Sprintf(CRedisErrors, errText)
}

/*RESPhandler interpret RESP in redis format
first char + may be +OK
- may be -Error
$ a good response from GET
: is an int response code */
func RESPhandler(resp []byte) (string, error) {
	for i, v := range resp {
		if i == 0 { //first char check
			switch v {
			case '+', ':', '-':
				return string(resp[1:]), nil
			case '$':
				return string(resp[2:]), nil
			}
		}
	}
	return "", fmt.Errorf("Redis Spec error in command")
}
