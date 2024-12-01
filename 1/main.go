package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	total := 0

	words := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for scanner.Scan() {
		textLn := scanner.Text()
		log.Printf("S: %v ", textLn)

		log.Printf("len: %v", len(textLn))

		first := strings.IndexAny(textLn, "0123456789")
		last := strings.LastIndexAny(textLn, "0123456789")

		firstNum, lastNum := "", ""
		if first != -1 {
			firstNum = string(textLn[first])
		}
		if last != -1 {
			lastNum = string(textLn[last])
		}

		for key, val := range words {
			temp := strings.Index(textLn, key)
			if temp != -1 && (temp < first || first == -1) {
				first = temp
				firstNum = string(strconv.Itoa(val)) //this is innefficient
			}
			if first == 0 {
				break
			}
		}

		for key, val := range words {
			temp := strings.LastIndex(textLn, key)
			if temp > last {
				last = temp
				lastNum = string(strconv.Itoa(val))
			}
			if last == len(textLn)-1 { //not sure if its -1 or -2
				break
			}
		}

		if first == -1 { //for now I dont think instructions say to ignore lines without numbers
			log.Fatal("No number found!")
		}

		log.Printf("I: %v,%v C: %v,%v", first, last, firstNum, lastNum)

		num, err := strconv.Atoi(string(firstNum + lastNum))

		if err != nil {
			log.Fatal(err)
		}

		log.Printf(" N:%v T:%v", num, total)
		log.Println()

		total += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total Amount: %v", total)
	log.Println()

}
