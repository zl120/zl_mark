package utils

import "encoding/json"

func DataToMap(data map[string]interface{}) (string,error) {
	strByte,err:=json.Marshal(data)
	if err != nil {
		return "",err
	}
	return string(strByte),nil
}
