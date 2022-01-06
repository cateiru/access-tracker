package tools

import (
	"bytes"
	"net/http"
)

// responseをstringに変換する
func ConvertResp(resp *http.Response) string {
	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	return buf.String()
}

// responseをbytesに変換する
func ConvertByteResp(resp *http.Response) []byte {
	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	return buf.Bytes()
}
