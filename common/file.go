package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	defer f.Close()
	names := make([]string, 0)
	if err != nil {
		return names, err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err == nil {
			name := strings.TrimSpace(line)
			if name != "" {
				names = append(names, name)
			}
		} else {
			if err == io.EOF {
				return names, nil
			}
			return names, err
		}
	}
	return names, nil
}

func WriteMapToFile(m map[string]int, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for k := range m {
		fmt.Fprintln(w, k)
	}
	return w.Flush()
}
