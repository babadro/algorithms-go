package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ParseNums[T int | int32 | int64 | float32 | float64](filePath string) []T {
	byteValue := bytesFromPath(filePath)
	var interfacesArr []interface{}
	if err := json.Unmarshal(byteValue, &interfacesArr); err != nil {
		panic(err)
	}

	res := make([]T, len(interfacesArr))
	for i := range interfacesArr {
		if interfacesArr[i] == nil {
			panic(fmt.Sprintf("nil value at index %d", i))
		}

		if value, ok := interfacesArr[i].(float64); ok {
			res[i] = T(value)
		} else {
			panic("Type assertion failed")
		}
	}

	return res
}

func ParseNullableNums[T int | int32 | int64 | float32 | float64](filePath string) []interface{} {
	byteValue := bytesFromPath(filePath)
	var result []interface{}
	if err := json.Unmarshal(byteValue, &result); err != nil {
		panic(err)
	}

	for i := range result {
		if result[i] == nil {
			continue
		}

		if value, ok := result[i].(float64); ok {
			result[i] = T(value)
		} else {
			panic("Type assertion failed")
		}
	}

	return result
}

func bytesFromPath(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	return byteValue
}
