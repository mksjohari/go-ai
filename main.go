package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func loading(c chan struct{}) {
	i := 0
	for {
		select {
		case <- c:
			fmt.Println()
			return
		default:
			fmt.Printf("\rLoading%s",strings.Repeat(".", i))
			i++
			if i > 3 {
				i = 0
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Print("\033[GLoading\033[K") // move the cursor left and clear the line
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

	for s.Scan() {
		dump := s.Text()
		fmt.Print("\rContinue? (y/n): ")
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

