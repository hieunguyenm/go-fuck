package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	tapeLength = flag.Int("length", 30000, "length of tape (can be expanded automatically)")
	pause      = flag.Int("pause", 200, "delay on console output in milliseconds")
	file       = flag.String("in", "", "path to Brainfuck file")
)

func main() {
	var fuck string
	var err error
	flag.Parse()
	fmt.Printf("Delay on console output is set to %dms.\n", *pause)

	if *file == "" {
		fmt.Println("Reading from stdin")
		fuck, err = bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalln("failed to read from stdin")
		}
	} else {
		f, err := os.Open(*file)
		if err != nil {
			log.Fatalln("failed to read source file provided")
		}
		defer f.Close()
		fuck, err = bufio.NewReader(f).ReadString('\n')
		if err != nil {
			log.Fatalln("failed to read Brainfuck in source file")
		}
	}

	if err = interpret(strings.TrimSuffix(fuck, "\n")); err != nil {
		fmt.Printf("failed to interpret: %v\n", err)
		os.Exit(1)
	}
}

func initTape(size int) []int {
	return make([]int, size)
}

func expand(tape []int) []int {
	t := tape[:int(float64(len(tape))*1.5)]
	return t
}

func interpret(fuck string) error {
	pointer := 0
	tape := initTape(*tapeLength)
	stdinReader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(fuck); i++ {
		if pointer < 0 {
			return fmt.Errorf("pointer is less than tape head index")
		}
		if pointer > len(tape)-1 {
			tape = expand(tape)
		}
		switch fuck[i] {
		case '.':
			fmt.Printf("%c", tape[pointer])
			time.Sleep(time.Duration(*pause) * time.Millisecond)
		case '>':
			pointer++
		case '<':
			pointer--
		case '+':
			tape[pointer]++
		case '-':
			tape[pointer]--
		case ',':
			b, err := stdinReader.ReadByte()
			if err != nil {
				return fmt.Errorf("failed to read byte from stdin")
			}
			tape[pointer] = int(b)
		case '[':
			if tape[pointer] == 0 {
				iter := 1
				for iter > 0 {
					i++
					if i > len(fuck)-1 {
						return fmt.Errorf("pointer out of bounds")
					}
					if fuck[i] == '[' {
						iter++
					}
					if fuck[i] == ']' {
						iter--
					}
				}
			}
		case ']':
			if tape[pointer] != 0 {
				iter := 1
				for iter > 0 {
					i--
					if i < 0 {
						return fmt.Errorf("pointer out of bounds")
					}
					if fuck[i] == ']' {
						iter++
					}
					if fuck[i] == '[' {
						iter--
					}
				}
			}
		default:
			return fmt.Errorf("invalid token at %d: %q", i, fuck[i])
		}
	}
	return nil
}
