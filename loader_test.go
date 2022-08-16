package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"

	"testing"

	"github.com/stretchr/testify/assert"
)

var sizes []string
var pointer int = 0
var mutex sync.Mutex

func FakeLoad(url *string) int {
	mutex.Lock()
	s, err := strconv.Atoi(sizes[pointer])
	if err != nil {
		panic(err.Error())
	}
	pointer++
	mutex.Unlock()
	return s
}

func TestLoader(t *testing.T) {

	for i := 0; i < 100; i++ {
		sizes = append(sizes, fmt.Sprintf("%d", rand.Int()))
	}

	err, records := Loader(sizes, FakeLoad)
	if err != nil {
		panic(err.Error())
	}

	sum1 := 0
	sum2 := 0
	for i := 0; i < len(records); i++ {
		s1, err := strconv.Atoi(sizes[i])
		if err != nil {
			panic(err)
		}
		s2, err := strconv.Atoi(records[i])
		if err != nil {
			panic(err)
		}
		sum1 += s1
		sum2 += s2
	}

	assert.Equal(t, sum1, sum2, "they should be equal")
}
