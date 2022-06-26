package _535_encode_and_decode_tinyURL

import "encoding/base64"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	return base64.StdEncoding.EncodeToString([]byte(longUrl))
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	res, _ := base64.StdEncoding.DecodeString(shortUrl)
	return string(res)
}
