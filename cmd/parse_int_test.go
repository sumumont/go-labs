package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
)

func TestParseInt(t *testing.T) {
	//fmt.Println(strconv.ParseInt("0lldnba41", 36, 10))
	filePath := "D:\\tmp\\data_channel.txt"
	dcAll := readLines(filePath)
	filePath1 := "D:\\tmp\\data_channel_j2.txt"
	j2All := readLines(filePath1)
	j2Map := map[string]struct{}{}
	for _, j2 := range j2All {
		j2Map[j2] = struct{}{}
	}

	fmt.Println("j2Map.len", len(j2Map))
	dcMap := map[string]string{}
	for _, str := range dcAll {
		strs := strings.Split(str, ",")
		if len(strs) <= 1 || len(strs[1]) == 0 {
			continue
		}
		uuid := strs[1]
		tags := strings.Split(uuid, "-")
		//fmt.Println(uuid,tags)
		//fmt.Println(tags)
		fileTime := tags[3]
		realTime, err := base36ToBase10(fileTime)
		if err != nil {
			panic(err)
		}
		dcMap[fmt.Sprintf("%d", realTime)] = fileTime
	}
	fmt.Println("dcMap.len", len(dcMap))
	//notFound := []string{}
	for k, _ := range j2Map {
		if _, ok := dcMap[k]; !ok {
			fmt.Println("this not upload infer: ", k)
		}
	}
}

func readLines(filePath string) []string {
	dcBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	dcStr := string(dcBytes)
	return strings.Split(dcStr, "\n")
}
func charToInt(char rune) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	}
	if char >= 'a' && char <= 'z' {
		return int(char-'a') + 10
	}
	if char >= 'A' && char <= 'Z' {
		return int(char-'A') + 10
	}
	return -1 // Invalid character
}

func base36ToBase10(input string) (int, error) {
	base := 36
	decimal := 0

	for i := len(input) - 1; i >= 0; i-- {
		charValue := charToInt(rune(input[i]))
		if charValue == -1 {
			return 0, fmt.Errorf("invalid character in input")
		}
		position := len(input) - 1 - i
		decimal += charValue * int(math.Pow(float64(base), float64(position)))
	}

	return decimal, nil
}
