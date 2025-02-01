package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func loading(c chan struct{}) {
	for {
		select {
		case <- c:
			fmt.Println()
			return
		default:
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func processText(text string, s *bufio.Scanner) bool {
	ch := make(chan struct{})
	go loading(ch)
	time.Sleep(2000 * time.Millisecond)
	ch <- struct{}{}
	close(ch)

	fmt.Println("<INSERT>\n")
	fmt.Println("Continue? (y/n)")

	for s.Scan() {
		dump := s.Text()
		if dump == "y" {
			return true
		} else if dump == "n" {
			return false
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Input: ")
		scanner.Scan()
		text := scanner.Text()
		another := processText(text, scanner)
		if !another {
			break
		}
	}
	fmt.Println("Thank you, Come Again.")
}

