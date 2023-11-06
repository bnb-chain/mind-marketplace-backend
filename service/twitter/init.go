package twitter

import (
	"github.com/bnb-chain/mind-marketplace-backend/util"
	tt "github.com/kkdai/twitter"
	"os"
)

var ConsumerKey string
var ConsumerSecret string
var CallbackURL string

func init() {
	//Twitter Dev Info from https://developer.twitter.com/en/apps
	ConsumerKey = os.Getenv("ConsumerKey")
	ConsumerSecret = os.Getenv("ConsumerSecret")

	// test purpose
	if ConsumerKey == "" {
		ConsumerKey = "XCEgb9jLLiByoG7cLldGbSmpj"
	}
	if ConsumerSecret == "" {
		ConsumerSecret = "TGzhpTxLlkUC8Yv572nPOMA4bVW8jMEt33mMJQQJLiaewVHkHL"
	}

	CallbackURL = os.Getenv("CallbackURL")

	// test purpose
	if CallbackURL == "" {
		CallbackURL = "http://www.localhost:8080/v1/twitter_token"
	}

	util.Logger.Info("Init server key=", ConsumerKey, " secret=", ConsumerSecret,
		" callback=", CallbackURL)
	twitterClient = tt.NewServerClient(ConsumerKey, ConsumerSecret)
}
