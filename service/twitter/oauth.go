package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bnb-chain/mind-marketplace-backend/service"
	"net/http"

	tt "github.com/kkdai/twitter"
)

var twitterClient *tt.ServerClient

func GetTwitterToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter Get twitter token")
	values := r.URL.Query()
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")
	address := values.Get("address")

	err := twitterClient.CompleteAuth(tokenKey, verificationCode)
	if err != nil {
		fmt.Println(err)
		errURL := fmt.Sprintf("https://%s/error?err=%s", r.Host, service.TwitterAuthErr)
		http.Redirect(w, r, errURL, http.StatusTemporaryRedirect)
		return
	}

	user, _, err := QueryUserProfile()
	if err != nil {
		fmt.Println(err)
		errURL := fmt.Sprintf("https://%s/error?err=%s", r.Host, service.TwitterAuthErr)
		http.Redirect(w, r, errURL, http.StatusTemporaryRedirect)
		return
	}

	fmt.Println(user.ScreenName)
	fmt.Println(address)

	_, err = service.AccountSvc.VerifyTwitter(context.Background(), address, user.ScreenName)
	if err != nil {
		fmt.Println(err)
		homeURL := fmt.Sprintf("https://%s/error?err=%s", r.Host, err)
		http.Redirect(w, r, homeURL, http.StatusTemporaryRedirect)
		return
	}

	homeURL := fmt.Sprintf("https://%s/", r.Host)
	http.Redirect(w, r, homeURL, http.StatusTemporaryRedirect)
}

func RedirectToTwitter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter redirect to twitter")
	fmt.Println("Token URL=", CallbackURL)

	values := r.URL.Query()
	address := values.Get("address")
	_, err := service.AccountSvc.Get(context.Background(), address)
	if err != nil {
		errorURL := fmt.Sprintf("https://%s/error?err=%s", r.Host, service.IncorrectAddressErr)
		http.Redirect(w, r, errorURL, http.StatusTemporaryRedirect)
		return
	}

	requestUrl, err := twitterClient.GetAuthURL(CallbackURL + "?address=" + address)
	if err != nil {
		errorURL := fmt.Sprintf("https://%s/error?err=%s", r.Host, service.CannotRedirectErr)
		http.Redirect(w, r, errorURL, http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, requestUrl, http.StatusTemporaryRedirect)
	fmt.Println("Leave redirect")
}

// User is the struct to fetch Twitter user name.
// https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/manage-account-settings/api-reference/get-account-settings
type User struct {
	ScreenName string `json:"screen_name"`
}

func QueryUserProfile() (User, []byte, error) {
	requestURL := fmt.Sprintf(tt.API_BASE+"%s", "account/settings.json")
	data, err := twitterClient.BasicQuery(requestURL)
	if err != nil {
		return User{}, nil, err
	}
	ret := User{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}
