package session

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// SaveCookies saves all browser cookies to a JSON file
func SaveCookies(browser *rod.Browser, path string) error {
	cookies, err := proto.NetworkGetAllCookies{}.Call(browser)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cookies.Cookies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// LoadCookies loads cookies from a JSON file into the browser
func LoadCookies(browser *rod.Browser, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var storedCookies []*proto.NetworkCookie
	if err := json.Unmarshal(data, &storedCookies); err != nil {
		return err
	}

	var cookieParams []*proto.NetworkCookieParam
	for _, c := range storedCookies {
		cookieParams = append(cookieParams, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}

	return proto.NetworkSetCookies{
		Cookies: cookieParams,
	}.Call(browser)
}


// CookiesExist checks whether the cookie file exists
func CookiesExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
