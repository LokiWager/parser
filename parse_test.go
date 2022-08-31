package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpaceShouldWork(t *testing.T) {
	word := " Hello World."
	tokens := Parse([]byte(word))
	assert.Equal(t, []string{"Hello", "World"}, tokens)
	assert.Equal(t, 2, len(tokens))
}

func TestNumberShouldWork(t *testing.T) {
	word := ` Hello World 2000-100.`
	tokens := Parse([]byte(word))
	assert.Equal(t, []string{"Hello", "World", "2000", "100"}, tokens)
	assert.Equal(t, 4, len(tokens))
}

func TestUnderlineShouldWork(t *testing.T) {
	word := ` Hello World in-
teresting.`
	tokens := Parse([]byte(word))
	assert.Equal(t, []string{"Hello", "World", "interesting"}, tokens)
	assert.Equal(t, 3, len(tokens))
}

func TestParseShouldWork(t *testing.T) {
	word := ` Hello World 100-20. in-
teresting 11.2 at 11:20 am`
	tokens := Parse([]byte(word))
	assert.Equal(t, []string{"Hello", "World", "100", "20", "interesting",
		"11.2", "at", "11:20 am"},
		tokens)
	assert.Equal(t, 8, len(tokens))
}
