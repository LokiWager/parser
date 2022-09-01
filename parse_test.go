package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpaceShouldWork(t *testing.T) {
	word := " Hello World.And"
	count := Parse([]byte(word))
	assert.Equal(t, 3, count)

	word = " Hello three-year-old boy ."
	count = Parse([]byte(word))
	assert.Equal(t, 3, count)

	word = "connect(Source.OutPort, FilterTransform.InPort)"
	assert.Panics(t, func() { Parse([]byte(word)) }, "Receive Error")

	word = `Hello
World`
	count = Parse([]byte(word))
	assert.Equal(t, 2, count)
}

func TestNumberShouldWork(t *testing.T) {
	word := ` Hello World 2000-100.`
	count := Parse([]byte(word))
	assert.Equal(t, 4, count)
}

func TestUnderlineShouldWork(t *testing.T) {
	word := ` Hello World in-
teresting.`
	count := Parse([]byte(word))
	assert.Equal(t, 3, count)
}

func TestParseShouldWork(t *testing.T) {
	word := ` Hello World 100-20. in-
teresting 11.2 at 11:20 am`
	count := Parse([]byte(word))
	assert.Equal(t, 9, count)
}
