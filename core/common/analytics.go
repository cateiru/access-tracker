package common

import (
	"encoding/json"

	ua "github.com/mileusna/useragent"
)

type UserAgent struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	OS        string `json:"os"`
	OSVersion string `json:"os_version"`
	Device    string `json:"device"`
	Mobile    bool   `json:"mobile"`
	Tablet    bool   `json:"access_id"`
	Desktop   bool   `json:"desktop"`
	Bot       bool   `json:"bot"`
	URL       string `json:"url"`
	String    string `json:"string"`
}

func NewUserAgent(target string) *UserAgent {
	parsed := ua.Parse(target)

	return &UserAgent{
		Name:      parsed.Name,
		Version:   parsed.Version,
		OS:        parsed.OS,
		OSVersion: parsed.OSVersion,
		Device:    parsed.Device,
		Mobile:    parsed.Mobile,
		Tablet:    parsed.Tablet,
		Desktop:   parsed.Desktop,
		Bot:       parsed.Bot,
		URL:       parsed.URL,
		String:    parsed.String,
	}
}

// userAgentを解析し、json形式で返します
func (c *UserAgent) ToJson() ([]byte, error) {
	return json.Marshal(c)
}

// UserAgentToJsonのjsonを構造体に戻します
func ParseUserAgentJson(target []byte) (*UserAgent, error) {
	var userAgent UserAgent

	if err := json.Unmarshal(target, &userAgent); err != nil {
		return nil, err
	}

	return &userAgent, nil
}
