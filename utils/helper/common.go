package helper

import (
	"crypto/rand"
	"encoding/json"
)

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap)
	return
}

func GenerateRandomNumber() uint {
	p, _ := rand.Prime(rand.Reader, 32)
	return uint(p.Uint64())
}
