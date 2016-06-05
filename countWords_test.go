package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNothingForEmptyString(t *testing.T) {
	assert.Equal(t, TotalWordCount(""), 0)
}

func TestShouldReturnNothingForWhitespace(t *testing.T) {
	assert.Equal(t, TotalWordCount("        "), 0)
}

func TestShouldCountOneWord(t *testing.T) {
	assert.Equal(t, SingleWordCount("foo", false),
		map[string]int{"foo": 1})
}

func TestShouldCountOneWordMultipleTimes(t *testing.T) {
	assert.Equal(t, SingleWordCount("foo foo", true),
		map[string]int{"foo": 2})
}

func TestShouldCountTwoWords(t *testing.T) {
	assert.Equal(t, SingleWordCount("foo bar", false),
		map[string]int{"foo": 1, "bar": 1})
}

func TestShouldCountMultipleWordsMultipleTimes(t *testing.T) {
	assert.Equal(t, SingleWordCount("foo bar yo foo yo test yarrr", true),
		map[string]int{
			"foo": 2, "bar": 1, "yo": 2,
			"test": 1, "yarrr": 1})
}
