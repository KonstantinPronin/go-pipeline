package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile(path string) (lines []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if ferr := file.Close(); ferr != nil {
			err = ferr
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeIntoFile(lines []string, path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if ferr := file.Close(); ferr != nil {
			err = ferr
		}
	}()

	w := io.Writer(file)

	for _, line := range lines {
		_, err = fmt.Fprintln(w, line)
		if err != nil {
			return err
		}
	}

	return nil
}
