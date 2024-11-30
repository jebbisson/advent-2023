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

	for scanner.Scan() {
		textLn := scanner.Text()
		log.Printf("S: %v ", textLn)

		log.Printf("len: %v", len(textLn))

		first := strings.IndexAny(textLn, "0123456789")
		last := strings.LastIndexAny(textLn, "0123456789")

		if first == -1 { //for now I dont think instructions say to ignore lines without numbers
			log.Fatal("No number found!")
		}

		//example show this as valid
		//if first == last {
		//	log.Fatal("First and last index are the same: %v %v", first, last)
		//}

		log.Printf("I: %v,%v", first, last)

		firstNum := string(textLn[first])
		lastNum := string(textLn[last])
		log.Printf("C: %v,%v", firstNum, lastNum)

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
