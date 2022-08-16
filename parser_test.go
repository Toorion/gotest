package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	urls := []string{"https://habr.com/ru/news/t/682652/",
		"https://habr.com/ru/news/t/682646/",
		"https://habr.com/ru/news/t/682640/",
		"https://habr.com/ru/news/t/682636/",
		"https://habr.com/ru/news/t/682634/",
		"https://habr.com/ru/news/t/682626/",
		"https://habr.com/ru/news/t/682616/",
		"https://habr.com/ru/news/t/682610/",
		"https://habr.com/ru/news/t/682608/",
		"https://habr.com/ru/news/t/682600/"}

	stringReader := strings.NewReader(strings.Join(urls, "\n"))
	rc := io.NopCloser(stringReader)
	err, lines := Parse(rc)
	if err != nil {
		panic(err.Error())
	}

	assert.Equal(t, len(urls), len(lines), "they should be equal")

	for i := 0; i < len(urls); i++ {
		assert.Equal(t, urls[i], lines[i], "they should be equal")
	}
}
