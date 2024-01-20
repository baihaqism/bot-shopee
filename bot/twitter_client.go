package bot

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/dghubble/oauth1"
)

type TwitterClient struct {
	api *anaconda.TwitterApi
}

func NewTwitterClient(config *Config) (*TwitterClient, error) {
	oauthConfig := &oauth1.Config{
		ConsumerKey:    config.Twitter.ConsumerKey,
		ConsumerSecret: config.Twitter.ConsumerSecret,
	}

	token := oauth1.NewToken(config.Twitter.AccessToken, config.Twitter.AccessTokenSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	api := anaconda.NewTwitterApi(httpClient)

	return &TwitterClient{api: api}, nil
}

func (c *TwitterClient) PostTweet(tweet string) error {
	_, err := c.api.PostTweet(tweet, nil)
	return err
}
