package json_utils

import "github.com/sirupsen/logrus"
import "encoding/json"

func IsJSON(str string) bool {
	var js interface{}
	return json.Unmarshal([]byte(str), &js) == nil
}

func ConvertToObject(jsonStr string, any any) error {
	err := json.Unmarshal([]byte(jsonStr), &any)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
func ConvertToString(any any) (string, error) {
	respStr, err := json.Marshal(any)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return string(respStr), nil
}
