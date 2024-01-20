package bot

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

type TweetCreator struct{}

func NewTweetCreator() *TweetCreator {
	return &TweetCreator{}
}

func (c *TweetCreator) CreateTweet(product *gofeed.Item, affiliateURL string) string {
	return fmt.Sprintf("%s\n\n%s", product.Title, affiliateURL)
}
