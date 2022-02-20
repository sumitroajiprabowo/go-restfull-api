package helper

import "strconv"

//create func with string to id with strconv.atoi
func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
