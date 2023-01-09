package url_generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const userId = "wqer7c6-7wer-qnof-fneiun"

func TestGenerateShortLink(t *testing.T) {
	longlink1 := "https://www.google.com"
	shortlink1 := GenerateShortUrl(longlink1, userId)

	longlink2 := "https://www.wikipedia.org"
	shortlink2 := GenerateShortUrl(longlink2, userId)

	assert.Equal(t, shortlink1, "RzXPyXYR")
	assert.Equal(t, shortlink2, "M85n9SAM")
}
