package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpaceShouldWork(t *testing.T) {
	word := " Hello World.And"
	count, _ := WordCount([]byte(word), true)
	assert.Equal(t, 3, count)

	word = `Hello
World`
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 2, count)

	word = ` : - 
Hello : 
World`
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 2, count)

	word = ` Hello World in-
teresting.`
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 3, count)

	word = "\tHello\rWorld in-\r\n\r\r\n\nternet"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 3, count)

	word = "\tHello\rWorld in-\r\n\r\r\n\n\tternet"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 4, count)

	word = "\tHello + World in-ternet * today is \\\\ Wed."
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 6, count)

	word = `The online encyclopedia project, Wikipedia, is the most popular wiki-
based website..`
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 11, count)
}

func TestDelimiterShouldWork(t *testing.T) {
	word := " Hello three-year-old boy ."
	count, _ := WordCount([]byte(word), true)
	assert.Equal(t, 3, count)

	word = "2000 - 100 = 1900"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 3, count)

	word = "11.11 12:00 13:00:01"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 3, count)

	word = "+86-9895-2700"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 1, count)

	word = "0001110651 Netease, Inc./ADR"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 4, count)

	word = "1.1.1.1 1.-89:01"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 2, count)

	word = "zoro201sword"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 3, count)
}

func TestNumberShouldWork(t *testing.T) {
	word := `100`
	count, _ := WordCount([]byte(word), true)
	assert.Equal(t, 1, count)

	word = "100\a200"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 2, count)

	word = `86-
2007`
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 2, count)

	word = "86-2007"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 1, count)
}

func TestErrorShouldWork(t *testing.T) {
	word := "中文 hello world"
	_, err := WordCount([]byte(word), true)
	assert.Equal(t, IllegalCharacter, err)

	word = "中文 hello world"
	count, err := WordCount([]byte(word), false)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, count)
}

func TestFinalStateShouldWork(t *testing.T) {
	word := "hello world\000"
	count, _ := WordCount([]byte(word), true)
	assert.Equal(t, 2, count)

	word = "hello\000 world\000"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 1, count)

	word = ""
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 0, count)

	word = "\000\000\000"
	count, _ = WordCount([]byte(word), true)
	assert.Equal(t, 0, count)
}
