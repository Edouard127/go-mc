package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const MojangAPI = "https://api.minecraftservices.com"

var DefaultProfile = Profile{"Steve", "5627dd98-e6be-3c21-b8a8-e92344183641"}
var DefaultAuth = Auth{DefaultProfile, Microsoft{}, KeyPair{}} // Offline-mode by default

type Profile struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}

type Auth struct {
	Profile
	Microsoft
	KeyPair
}

func (a *Auth) createRequest(method, path string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, MojangAPI+path, body)
	if err != nil {
		panic(err)
	}

	if _, ok := body.(*strings.Reader); ok {
		req.Header.Set("Content-Type", "application/json")
	}

	if a.Microsoft.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+a.Microsoft.AccessToken)
	}

	return req
}

func (a *Auth) NameAvailable(name string) (bool, error) {
	resp, err := http.DefaultClient.Do(a.createRequest("GET", "/minecraft/profile/name/"+name+"/available", nil))
	if err != nil {
		return false, err
	}

	if resp.StatusCode == 401 {
		return false, fmt.Errorf("invalid access token")
	}

	var response struct {
		Status string `json:"status"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)

	return response.Status == "AVAILABLE", err
}

func (a *Auth) ChangeName(name string) error {
	resp, err := http.DefaultClient.Do(a.createRequest("PUT", "/minecraft/profile/name/"+name, nil))
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 401:
		return fmt.Errorf("invalid access token")
	case 403:
		return fmt.Errorf("name is unavailable")
	case 429:
		return fmt.Errorf("too many requests")
	case 500:
		return fmt.Errorf("internal server error")
	default:
		return nil
	}
}

func (a *Auth) ChangeSkin(variant, skinURL string) error {
	resp, err := http.DefaultClient.Do(a.createRequest("POST", "/minecraft/profile/skins", strings.NewReader(fmt.Sprintf(`{"variant": "%s"', "url":"%s"}`, variant, skinURL))))
	if err != nil {
		return err
	}

	if resp.StatusCode == 401 {
		return fmt.Errorf("invalid access token")
	}

	return nil
}

func (a *Auth) ResetSkin() error {
	resp, err := http.DefaultClient.Do(a.createRequest("DELETE", "/minecraft/profile/skins/active", nil))
	if err != nil {
		return err
	}

	if resp.StatusCode == 401 {
		return fmt.Errorf("invalid access token")
	}

	return nil
}

func (a *Auth) HideCape() error {
	resp, err := http.DefaultClient.Do(a.createRequest("DELETE", "/minecraft/profile/capes/active", nil))
	if err != nil {
		return err
	}

	if resp.StatusCode == 401 {
		return fmt.Errorf("invalid access token")
	}

	return nil
}

func (a *Auth) ShowCape(capeid string) error {
	resp, err := http.DefaultClient.Do(a.createRequest("POST", "/minecraft/profile/capes", strings.NewReader(fmt.Sprintf(`{"capeId":"%s"}`, capeid))))
	if err != nil {
		return err
	}

	if resp.StatusCode == 401 {
		return fmt.Errorf("invalid access token")
	}

	return nil
}

func (a *Auth) SessionID() string {
	return fmt.Sprintf("token:%s:%s", a.Microsoft.AccessToken, a.Profile.UUID)
}

type Microsoft struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	ExpiresAt    int64  `json:"expires_at"`
}

type DeviceCodeRequest struct {
	UserCode        string `json:"user_code"`
	DeviceCode      string `json:"device_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       int64  `json:"expires_in"`
	Interval        int64  `json:"interval"`
	Message         string `json:"message"`
}

type DeviceCodeResponse struct {
	Microsoft
	Error string `json:"error"`
}

type XboxLiveAuth struct {
	DisplayClaims struct {
		Xui []struct {
			Uhs string `json:"uhs"`
		} `json:"xui"`
	} `json:"DisplayClaims"`
	IssueInstant string `json:"IssueInstant"`
	NotAfter     string `json:"NotAfter"`
	Token        string `json:"Token"`
}

func PreAuthData(r io.Reader) (string, string, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return "", "", err
	}

	return Match(`sFTTag:'.*value=\"([^\"]*)\"`, string(buf)), Match(`urlPost:?'(.+?('))`, string(buf)), nil
}

func ExtractValue(url, key string) string {
	return strings.Split(url, key+"=")[1]
}

func Match(regex, content string) string {
	return regexp.MustCompile(regex).FindStringSubmatch(content)[1]
}
