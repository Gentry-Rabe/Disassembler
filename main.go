package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// removes spaces from integer string, checks if string contains only 1s and 0s
func cleanString(input string) string {
	grab := []rune(input)
	var runes []rune
	for i := 0; i < len(grab); i++ {
		if grab[i] == '1' || grab[i] == '0' {
			runes = append(runes, grab[i])
		}
	}
	if len(runes) == 32 {
		return string(runes)
	} else {
		return "Error: bad string length"
	}
}

// converts binary to decimal
func binToDec(input string) int {
	chars := strings.SplitAfter(input, "")
	if len(chars) == 0 {
		return 0
	}
	firstNum, err := strconv.Atoi(chars[0])
	if err != nil {
		log.Fatalf("Failed to convert binary character %s", err)
	}

	return firstNum*int(math.Pow(2, float64(len(chars)-1))) + binToDec(strings.Join(chars[1:], ""))

}

// gets opCode from binary
func getOpCode(input string, count int) {
	/*
		op codes are parsed from biggest to smallest.
		find 6-digit codes first, then 8, then 9, then 10, then 11.
		the switch cases are nested in the previous cases default,
		so that if the code doesn't find any opCodes of length X,
		it moves on to the next highest length.
	*/
	//6-digit opCode
	opcode := binToDec(input[0:11])

	switch {
	case opcode >= 160 && opcode <= 191:
		fmt.Print(input[0:6], " ", input[6:32], " ", count, " B\n")

	case opcode >= 1440 && opcode <= 1447:
		fmt.Print(input[0:8], " ", input[8:27], " ", input[27:32], " ", count, " CBZ\n")

	case opcode >= 1448 && opcode <= 1455:
		fmt.Print(input[0:8], " ", input[8:27], " ", input[27:32], " ", count, " CBNZ\n")

	case opcode >= 1684 && opcode <= 1687:
		fmt.Print(input[0:9], " ", input[9:11], " ", input[11:27], " ", input[27:32], " ", count, " MOVZ\n")

	case opcode >= 1940 && opcode <= 1943:
		fmt.Print(input[0:9], " ", input[9:11], " ", input[11:27], " ", input[27:32], " ", count, " MOVK\n")

	case opcode >= 1160 && opcode <= 1161:
		fmt.Print(input[0:10], " ", input[10:22], " ", input[22:27], " ", input[27:32], " ", count, " ADDI\n")

	case opcode >= 1672 && opcode <= 1673:
		fmt.Print(input[0:10], " ", input[10:22], " ", input[22:27], " ", input[27:32], " ", count, " SUBI\n")

	case opcode == 1104:
		fmt.Print(input[0:11], " ", input[11:16], " ", input[16:22], " ", input[22:27], " ", input[27:32], " ", count, " AND\n")

	case opcode == 1112:
		fmt.Print(input[0:11], " ", input[11:16], " ", input[16:22], " ", input[22:27], " ", input[27:32], " ", count, " ADD\n")

	case opcode == 1360:
		fmt.Print(input[0:11], " ", input[11:16], " ", input[16:22], " ", input[22:27], " ", input[27:32], " ", count, " ORR\n")

	case opcode == 1624:
		fmt.Print(input[0:11], " ", input[11:16], " ", input[16:22], " ", input[22:27], " ", input[27:32], " ", count, " SUB\n")

	case opcode == 1690:
		fmt.Print(input[0:11], " ", input[11:16], " ", input[16:22], " ", input[22:27], " ", input[27:32], " ", count, " LSR\n")

	case opcode == 1691:
		fmt.Print(input[0:11], " ", input[11:16], " ", input[16:22], " ", input[22:27], " ", input[27:32], " ", count, " LSL\n")

	case opcode == 1984:
		fmt.Print(input[0:11], " ", input[11:20], " ", input[20:22], " ", input[22:27], " ", input[27:32], " ", count, " STUR\n")

	case opcode == 1986:
		fmt.Print(input[0:11], " ", input[11:20], " ", input[20:22], " ", input[22:27], " ", input[27:32], " ", count, " LDUR\n")

	case opcode == 1692:
		fmt.Print(input[0:11], " ", input[11:16], " ", input[16:22], " ", input[22:27], " ", input[27:32], " ", count, " ASR\n")

	case opcode == 1872:
		fmt.Print(input[0:11], " ", input[11:16], " ", input[16:22], " ", input[22:27], " ", input[27:32], " ", count, " EOR\n")

	case opcode == 2038:
		fmt.Print(input, " ", count, " BREAK\n")

	case opcode == 0:
		fmt.Print("NOP")

	default:
		fmt.Print(input, " opCode not found\n")
	}
}

func main() {
	//grab file name pointers
	var InputFileName *string = flag.String("i", "", "Gets the input file name")
	var OutputFileName *string = flag.String("o", "", "Gets the output file name")
	flag.Parse()

	//document process in output, removable
	fmt.Print("Input: ", *InputFileName, "\n")
	fmt.Print("Output: ", *OutputFileName, "\n")

	if flag.NArg() != 0 {
		os.Exit(200)
	}

	//scan by opening file from pointer
	file, err := os.Open(*InputFileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	//scan
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	closeErr := file.Close()
	if closeErr != nil {
		log.Fatalf("Failed to close file: %s", closeErr)
	}

	//output (eventually for disassem. code)
	for i := range txtlines {
		iString := cleanString(txtlines[i])
		getOpCode(iString, 96+(i*4))
	}
}
