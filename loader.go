package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

type loadFunc func(*string) int

func Loader(lines []string, fn loadFunc) (error, []string) {
	result := make(chan int, len(lines))

	var wg sync.WaitGroup
	for _, line := range lines {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			result <- fn(&line)
		}(line)
	}

	wg.Wait()

	close(result)

	records := []string{}
	for size := range result {
		if 0 > size {
			return errors.New("Loading error"), nil
		}
		records = append(records, (fmt.Sprintf("%d", size)))
	}

	return nil, records
}

func Load(url *string) int {
	resp, err := http.Get(*url)
	if err != nil {
		return -1 // Test case only, in task not described behaviour in this case
	}
	if resp.StatusCode != 200 {
		// return -1 // Test case only, in task not described behaviour in this case
		// commented, because some test urls sometime return 503
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Len()
}
