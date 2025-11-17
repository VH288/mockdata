// Package data
package data

import (
	"fmt"
	"math/rand"
	"strings"
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
	year := 1950 + rand.Intn(100)
	month := 1 + rand.Intn(12)
	day := 1 + rand.Intn(28)

	return fmt.Sprintf("%02d-%02d-%d", day, month, year)
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
	prefixLen := 6 + rand.Intn(4)

	var sb strings.Builder
	sb.WriteString("081")

	for range prefixLen {
		digit := rand.Intn(10)
		digitString := fmt.Sprintf("%d", digit)

		sb.WriteString(digitString)
	}

	result := sb.String()
	return result
}
