package url_generator

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
)

func urlToHash(url string) []byte {
	encoder := sha256.New()
	encoder.Write([]byte(url))
	return encoder.Sum(nil)
}

func hashToText(hash []byte) string {
	converter := base58.BitcoinEncoding
	converted_bytes, err := converter.Encode(hash)
	if err != nil {
		panic(fmt.Sprintf("Cannot encode hash to text with error: %v", err))
	}
	converted_text := string(converted_bytes)
	return string(converted_text)
}

func GenerateShortUrl(longUrl string, userId string) string {
	hashedLongLink := urlToHash(longUrl + userId)
	generatedNumber := new(big.Int).SetBytes(hashedLongLink).Uint64()
	shortUrl := hashToText([]byte(fmt.Sprintf("%d", generatedNumber)))
	return shortUrl[:8]
}
