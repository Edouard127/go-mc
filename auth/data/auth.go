package data

import (
	"io"
	"regexp"
	"strings"
)

var DefaultProfile = Profile{"Steve", "5627dd98-e6be-3c21-b8a8-e92344183641"}

type Profile struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}

type Auth struct {
	Profile
	Microsoft
	KeyPair
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
