package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]
	f, _ := os.OpenFile("output.hex", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// fmt.Println(f.Name())

	dat, _ := os.ReadFile(argsWithoutProg[0])
	lines := strings.Split(string(dat), "\n")
	bin := ""
	hex := ""
	for i := 0; i < len(lines); i++ {
		binTemp, hexTemp := compileLine(lines[i])
		bin += binTemp
		hex += hexTemp
	}

	// fmt.Println(bin)
	// fmt.Println(hex)
	os.Truncate(f.Name(), 0)
	f.Write([]byte(hex))
}

func compileLine(line string) (string, string) {
	parts := strings.Fields(line)
	bin := mapOpCode(parts[0])

	switch parts[0] {
	//RRR type
	case "add", "nand":
		bin += mapReg(parts[1])
		bin += mapReg(parts[2])
		bin += "000"
		bin += mapReg(parts[3])
	//RRI type
	case "addi", "lw", "sw", "beq":
		bin += mapReg(parts[1])
		bin += mapReg(parts[2])
		bin += mapSigned6bit(parts[3])
	//RI type
	case "lui":
		bin += mapReg(parts[1])
		bin += mapUnsigned10bit(parts[2])
	case "jalr":
		bin += mapReg(parts[1])
		bin += mapReg(parts[2])
		bin += "0000000"
	}
	return bin, binaryToHex(bin) + " "
}

func mapOpCode(op string) string {
	switch op {
	case "add":
		return "000"
	case "addi":
		return "001"
	case "nand":
		return "010"
	case "lw":
		return "100"
	case "sw":
		return "101"
	case "beq":
		return "110"
	case "jalr":
		return "111"
	case "lui":
		return "011"
	}
	return ""
}
func mapReg(reg string) string {
	switch reg {
	case "r0":
		return "000"
	case "r1":
		return "001"
	case "r2":
		return "010"
	case "r3":
		return "011"
	case "r4":
		return "100"
	case "r5":
		return "101"
	case "r6":
		return "110"
	case "r7":
		return "111"
	}
	return ""
}
func mapSigned6bit(num string) string {
	if num == "" {
		return "000000"
	}
	binary := ""
	if num[0] == '-' {
		binary += "1"
	} else {
		binary += "0"
	}
	//convert to binary
	num = strings.Replace(num, "-", "", -1)
	numInt, _ := strconv.Atoi(num)
	binary += fmt.Sprintf("%06b", numInt)
	return binary
}
func mapUnsigned10bit(num string) string {
	//convert to binary
	numInt, _ := strconv.Atoi(num)
	return fmt.Sprintf("%010b", numInt)
}
func binaryToHex(binary string) string {
	//convert to hex
	numInt, _ := strconv.ParseInt(binary, 2, 64)
	return strconv.FormatInt(numInt, 16)
}
