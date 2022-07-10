package base

import (
	"net/http"

	"github.com/cateiru/go-client-hints/ch"
	"github.com/cateiru/go-client-hints/ch/headers"
	ua "github.com/mileusna/useragent"
)

var browsers = map[string]string{
	"Google Chrome":  "Chrome",
	"Microsoft Edge": "Edge",
}

type UserData struct {
	Device  string
	OS      string
	Browser string
}

func ParseUA(r *http.Request) *UserData {
	platform := r.Header.Get(headers.SecChUaPlatform)

	// UA-CHが存在しない場合はUser-Agentを使用してユーザ情報を取得する
	if platform == "" {
		return ParseUserAgent(r)
	}

	userData := ch.ParseSecChUa(r.Header.Get(headers.SecChUa))
	b := "Unknown"

	for browser, bName := range browsers {
		version := userData[browser]
		if version != "" {
			b = bName
			break
		}
	}

	deviceType := ch.ParseChUaMobile(r.Header.Get(headers.SecChUaMobile))
	device := "Unknown"

	switch deviceType {
	case ch.Mobile:
		device = "Mobile"
	case ch.NoMobile:
		device = "Desktop"
	}

	return &UserData{
		Device:  device,
		Browser: b,
		OS:      platform[1 : len(platform)-1],
	}
}

func ParseUserAgent(r *http.Request) *UserData {
	ua := ua.Parse(r.Header.Get("User-Agent"))

	device := ua.Device

	if device == "" {
		device = "Desktop"
	}

	return &UserData{
		Device:  device,
		OS:      ua.OS,
		Browser: ua.Name,
	}
}
