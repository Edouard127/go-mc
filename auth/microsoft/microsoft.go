package microsoft

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Edouard127/go-mc/auth/data"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"
)

var cookieJar, _ = cookiejar.New(nil)
var client = &http.Client{
	Jar: cookieJar,
}
var loginParams = map[string]string{
	"client_id":     max(XboxLiveClientID, os.Getenv("AZURE_CLIENT_ID")),
	"redirect_uri":  LiveRedirectURI,
	"scope":         XboxLiveScope,
	"response_type": "token",
	"display":       "touch",
	"locale":        "en",
}

func createRequest(at string, method string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, at, body)
	if err != nil {
		panic(err)
	}
	return req
}

func get(at string, headers map[string]string) (*http.Response, error) {
	req := createRequest(at, http.MethodGet, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return client.Do(req)
}

func postQuery(at string, data map[string]string) (*http.Response, error) {
	return client.Do(createRequest(at, http.MethodPost, strings.NewReader(urlValues(data))))
}

func postForm(at string, data map[string]any) (*http.Response, error) {
	b, _ := json.Marshal(data)
	req := createRequest(at, http.MethodPost, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return client.Do(req)
}

// LoginWithCredentials logs in with an email and a password.
// Not working
func LoginWithCredentials(username, password string) (data.XboxLiveAuth, error) {
	resp, err := get(MicrosoftAuthorizationEndpoint+"?"+urlValues(loginParams), nil)
	if err != nil {
		return data.XboxLiveAuth{}, err
	}

	ppft, postUrl, err := data.PreAuthData(resp.Body)
	if err != nil {
		return data.XboxLiveAuth{}, err
	}

	resp, err = postForm(postUrl[:len(postUrl)-1], map[string]any{
		"login":    username,
		"loginfmt": username,
		"passwd":   password,
		"PPFT":     ppft,
	})

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data.XboxLiveAuth{}, err
	}

	fmt.Println(string(bytes))

	loc, err := resp.Location()
	if err != nil {
		fmt.Println(err)
		return data.XboxLiveAuth{}, err
	}

	temp, err := getTokens(loc.String())
	if err != nil {
		return data.XboxLiveAuth{}, err
	}

	fmt.Println(temp)

	return xboxLiveLogin(temp.AccessToken)
}

func LoginWithDeviceCode() (data.XboxLiveAuth, error) {
	resp, err := postQuery(MicrosoftDeviceCodeEndpoint, map[string]string{
		"client_id": MicrosoftClientID,
		"scope":     MicrosoftScope,
	})
	if err != nil {
		return data.XboxLiveAuth{}, err
	}

	var deviceCodeReq data.DeviceCodeRequest
	json.NewDecoder(resp.Body).Decode(&deviceCodeReq)
	fmt.Printf("Please go to %s and enter %s to authenticate.\n", deviceCodeReq.VerificationURI, deviceCodeReq.UserCode)

	var deviceCodeResp data.DeviceCodeResponse

	for {
		resp, err = postQuery(MicrosoftTokenEndpoint, map[string]string{
			"client_id":   MicrosoftClientID,
			"scope":       MicrosoftScope,
			"grant_type":  "urn:ietf:params:oauth:grant-type:device_code",
			"device_code": deviceCodeReq.DeviceCode,
		})
		if err != nil {
			return data.XboxLiveAuth{}, err
		}

		json.NewDecoder(resp.Body).Decode(&deviceCodeResp)

		if resp.StatusCode == 400 {
			switch deviceCodeResp.Error {
			case "authorization_pending":
				continue
			case "slow_down":
				time.Sleep(time.Duration(deviceCodeReq.Interval) * time.Second)
				continue
			case "expired_token":
				return data.XboxLiveAuth{}, fmt.Errorf("device code expired")
			case "invalid_grant":
				return data.XboxLiveAuth{}, fmt.Errorf("invalid device code")
			}
		}

		if resp.StatusCode == 200 {
			return xboxLiveLogin(deviceCodeResp.AccessToken)
		}

		time.Sleep(time.Duration(deviceCodeReq.Interval) * time.Second)
	}
}

func xboxLiveLogin(accessToken string) (data.XboxLiveAuth, error) {
	resp, err := postForm(XboxLiveAuthorizationEndpoint, map[string]any{
		"Properties": map[string]any{
			"AuthMethod": "RPS",
			"SiteName":   XboxLiveAuthHost,
			"RpsTicket":  "d=" + accessToken,
		},
		"RelyingParty": XboxLiveAuthRelay,
		"TokenType":    "JWT",
	})
	if err != nil {
		return data.XboxLiveAuth{}, err
	}

	var xboxResp data.XboxLiveAuth
	json.NewDecoder(resp.Body).Decode(&xboxResp)

	return xboxSecurityLogin(xboxResp)
}

func xboxSecurityLogin(xboxResp data.XboxLiveAuth) (data.XboxLiveAuth, error) {
	resp, err := postForm(XboxXSTSAuthorizationEndpoint, map[string]any{
		"Properties": map[string]any{
			"SandboxId":  "RETAIL",
			"UserTokens": []string{xboxResp.Token},
		},
		"RelyingParty": MinecraftAuthRelay,
		"TokenType":    "JWT",
	})
	if err != nil {
		return data.XboxLiveAuth{}, err
	}

	var xboxSecure data.XboxLiveAuth
	json.NewDecoder(resp.Body).Decode(&xboxSecure)

	return xboxSecure, nil
}

// MinecraftLogin logs in with a data.XboxLiveAuth and returns a Minecraft account.
// If save is true, the account will be saved to the accounts.json file.
func MinecraftLogin(xboxSecure data.XboxLiveAuth, save bool) (data.Auth, error) {
	resp, err := postForm(MinecraftAuthorizationEndpoint, map[string]any{
		"identityToken": "XBL3.0 x=" + xboxSecure.DisplayClaims.Xui[0].Uhs + ";" + xboxSecure.Token,
	})
	if err != nil {
		return data.Auth{}, err
	}

	var minecraftResp data.Auth
	json.NewDecoder(resp.Body).Decode(&minecraftResp)
	minecraftResp.MicrosoftAuth.ExpiresAt = time.Now().Unix() + minecraftResp.ExpiresIn
	minecraftResp.Profile, err = MinecraftProfile(minecraftResp)

	if save {
		err = WriteMinecraftAccount(minecraftResp)
	}

	return minecraftResp, err
}

// MinecraftRefresh refreshes the access token of a Minecraft account.
// Can be used for login as well. If the account has no profile, it will
// be fetched.
func MinecraftRefresh(auth data.Auth) (data.Auth, error) {
	if auth.MicrosoftAuth.ExpiresAt > time.Now().Unix() {
		return auth, nil
	}

	resp, err := postForm(LiveTokenRefresh, map[string]any{
		"client_id":     MicrosoftClientID,
		"refresh_token": auth.RefreshToken,
		"grant_type":    "refresh_token",
		"redirect_uri":  MicrosoftNativeClient,
	})
	if err != nil {
		return data.Auth{}, err
	}

	var authResp data.MicrosoftAuth
	json.NewDecoder(resp.Body).Decode(&authResp)

	authResp.ExpiresAt = time.Now().Unix() + authResp.ExpiresIn
	auth.MicrosoftAuth = authResp

	if auth.Profile.UUID == "" {
		auth.Profile, err = MinecraftProfile(auth)
		if err != nil {
			return data.Auth{}, err
		}
	}

	return auth, nil
}

// MinecraftProfile fetches the profile of a Minecraft account.
func MinecraftProfile(auth data.Auth) (data.Profile, error) {
	resp, err := get(MinecraftProfileEndpoint, map[string]string{
		"Authorization": "Bearer " + auth.AccessToken,
	})
	if err != nil {
		return data.Profile{}, err
	}

	var profile data.Profile
	json.NewDecoder(resp.Body).Decode(&profile)

	return profile, nil
}

func WriteMinecraftAccount(account data.Auth) error {
	f := GetAccountFile()
	defer f.Close()

	var accounts []data.Auth
	accounts = append(accounts, account)
	fmt.Println(accounts)

	json.NewEncoder(f).Encode(accounts)
	fmt.Println(accounts)

	return nil
}

func LoginFromCache(f func(auth data.Auth) bool) data.Auth {
	accounts, err := ReadMinecraftAccounts()
	if err != nil {
		return data.Auth{}
	}

	if f == nil {
		return accounts[0]
	}

	for _, account := range accounts {
		if f(account) {
			return account
		}
	}

	return data.Auth{}
}

func ReadMinecraftAccounts() ([]data.Auth, error) {
	var accounts []data.Auth

	f := GetAccountFile()
	defer f.Close()

	json.NewDecoder(f).Decode(&accounts)

	if len(accounts) == 0 {
		return accounts, fmt.Errorf("no accounts found")
	}

	var err error

	for i := range accounts {
		accounts[i], err = MinecraftRefresh(accounts[i])
		if err != nil {
			fmt.Println("Error refreshing account: ", err)
		}
	}

	return accounts, nil
}

func GetAccountFile() *os.File {
	dir, _ := os.UserConfigDir()
	f, _ := os.OpenFile(fmt.Sprintf("%s/.go-mc/accounts.json", dir), os.O_CREATE|os.O_RDWR, 0644)
	return f
}

func getTokens(url string) (data.MicrosoftAuth, error) {
	var msauth data.MicrosoftAuth

	msauth.AccessToken = data.Match(`access_token=([^&]+)`, url)
	msauth.RefreshToken = data.Match(`refresh_token=([^&]+)`, url)
	return msauth, nil
}

func urlValues(m map[string]string) string {
	var sb strings.Builder
	for k, v := range m {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(v)
		sb.WriteString("&")
	}
	return sb.String()
}
