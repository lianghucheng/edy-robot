package conf

import (
	"bufio"
	"github.com/name5566/leaf/log"
	"io"
	"os"
	"strings"
)

func ReadName(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	defer f.Close()
	log.Debug("f: %v", fileName)
	var names []string
	if err != nil {
		return names, err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err == nil {
			names = append(names, strings.TrimSpace(line))
		} else {
			if err == io.EOF {
				return names, nil
			}
			return names, err
		}

	}
	return names, nil
}
