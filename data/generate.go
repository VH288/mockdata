// Package data
package data

import (
	"fmt"
	"math/rand"
)

func Generate(dataType string) any {
	switch dataType {
	case TypeName:
		return generateName()
	case TypeDate:
		return generateDate()
	case TypeAddress:
		return generateAddress()
	case TypePhone:
		return generatePhone()
	}
	return fmt.Sprintf("%s mock", dataType)
}

func generateName() string {
	nameLen := len(name)

	index := rand.Intn(nameLen)

	return name[index]
}

func generateDate() string {
	return ""
}

func generateAddress() string {
	streetLen := len(address[SubTypeStreet])
	cityLen := len(address[SubTypeCity])

	streetIndex := rand.Intn(streetLen)
	cityIndex := rand.Intn(cityLen)
	number := rand.Intn(100)

	return fmt.Sprintf("Jl. %s No. %d, %s", address[SubTypeStreet][streetIndex], number, address[SubTypeCity][cityIndex])
}

func generatePhone() string {
	return ""
}
