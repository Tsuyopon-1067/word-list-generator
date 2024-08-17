package scraping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrapingOk(t *testing.T) {
	word := "scraping"

	expected := "こすること、削ること、削り落としたもの、かきくず"
	actual := FetchWordMeaning(word)

	assert.Equal(t, expected, actual)
}
