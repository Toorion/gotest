package main

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func Parse(body io.ReadCloser) (error, []string) {

	lines := []string{}

	rd := bufio.NewReader(body)
	for {
		line, err := rd.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			//io.WriteString(w, "Read Error: "+err.Error()+"\n")
			return errors.New("Parse error"), nil
		}
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
	}

	if len(lines) > MAX_LINE_PER_REQUEST {
		return errors.New("Records too much"), nil
	}

	return nil, lines
}
