package base_test

import (
	"net/http"
	"testing"

	"github.com/cateiru/access-tracker/base"
	"github.com/stretchr/testify/require"
)

func TestParseUserAgent(t *testing.T) {
	t.Run("macのchromeで正しい", func(t *testing.T) {
		r := &http.Request{
			Header: map[string][]string{
				"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36"},
			},
		}

		user := base.ParseUserAgent(r)

		require.Equal(t, user, &base.UserData{
			Device:  "Desktop",
			OS:      "macOS",
			Browser: "Chrome",
		})
	})

	t.Run("windowsのedgeで正しい", func(t *testing.T) {
		r := &http.Request{
			Header: map[string][]string{
				"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36 Edg/94.0.992.31"},
			},
		}

		user := base.ParseUserAgent(r)

		require.Equal(t, user, &base.UserData{
			Device:  "Desktop",
			OS:      "Windows",
			Browser: "Edge",
		})
	})

	t.Run("windowsのfirefoxで正しい", func(t *testing.T) {
		r := &http.Request{
			Header: map[string][]string{
				"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0"},
			},
		}

		user := base.ParseUserAgent(r)

		require.Equal(t, user, &base.UserData{
			Device:  "Desktop",
			OS:      "Windows",
			Browser: "Firefox",
		})
	})

	t.Run("macのsafariで正しい", func(t *testing.T) {
		r := &http.Request{
			Header: map[string][]string{
				"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Safari/605.1.15"},
			},
		}

		user := base.ParseUserAgent(r)

		require.Equal(t, user, &base.UserData{
			Device:  "Desktop",
			OS:      "macOS",
			Browser: "Safari",
		})
	})

	t.Run("iosのsafariで正しい", func(t *testing.T) {
		r := &http.Request{
			Header: map[string][]string{
				"User-Agent": {"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1"},
			},
		}

		user := base.ParseUserAgent(r)

		require.Equal(t, user, &base.UserData{
			Device:  "iPhone",
			OS:      "iOS",
			Browser: "Safari",
		})
	})
}

func TestPraseUA(t *testing.T) {
	t.Run("chromeでclient hintsが存在する場合はそれを使用する", func(t *testing.T) {
		r := &http.Request{
			Header: map[string][]string{
				"Sec-Ch-Ua-Platform": {`"Windows"`},
				"Sec-Ch-Ua":          {`" Not A;Brand";v="99", "Chromium";v="96", "Google Chrome";v="96"`},
				"Sec-Ch-Ua-Mobile":   {"?0"},
			},
		}

		user := base.ParseUA(r)

		require.Equal(t, user, &base.UserData{
			Device:  "Desktop",
			OS:      "Windows",
			Browser: "Chrome",
		})
	})

	t.Run("chromeでモバイル端末のときでclient hintsを利用する", func(t *testing.T) {
		r := &http.Request{
			Header: map[string][]string{
				"Sec-Ch-Ua-Platform": {`"Android"`},
				"Sec-Ch-Ua":          {`" Not A;Brand";v="99", "Chromium";v="96", "Google Chrome";v="96"`},
				"Sec-Ch-Ua-Mobile":   {"?1"},
			},
		}

		user := base.ParseUA(r)

		require.Equal(t, user, &base.UserData{
			Device:  "Mobile",
			OS:      "Android",
			Browser: "Chrome",
		})
	})

	t.Run("client hintsがない場合はuaを使用する", func(t *testing.T) {
		r := &http.Request{
			Header: map[string][]string{
				"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36"},
			},
		}

		user := base.ParseUA(r)

		require.Equal(t, user, &base.UserData{
			Device:  "Desktop",
			OS:      "macOS",
			Browser: "Chrome",
		})
	})
}
