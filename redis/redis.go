package redis

import "fmt"

/*CRedisFormat holds the redis communication format*/
const CRedisFormat = "$%d\r\n%s\r\n"

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
