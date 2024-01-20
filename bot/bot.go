package bot

import (
	"log"
	"time"
)

type Bot struct {
	twitterClient *TwitterClient
	shopeeClient  *ShopeeClient
	tweetCreator  *TweetCreator
}

func NewBot() (*Bot, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	twitterClient, err := NewTwitterClient(config)
	if err != nil {
		return nil, err
	}

	shopeeClient := NewShopeeClient(config)
	tweetCreator := NewTweetCreator()

	return &Bot{
		twitterClient: twitterClient,
		shopeeClient:  shopeeClient,
		tweetCreator:  tweetCreator,
	}, nil
}

func (b *Bot) Start() error {
	for {
		products, err := getLatestProducts("https://shopee.com.my/api/v4/feeds/feed_item_list?limit=10&feed_type=1&official_shop=1&sort_type=3")
		if err != nil {
			log.Println(err)
			continue
		}

		for _, product := range products {
			affiliateURL := b.shopeeClient.CreateAffiliateURL(product.Link)
			tweet := b.tweetCreator.CreateTweet(product, affiliateURL)
			if err := b.twitterClient.PostTweet(tweet); err != nil {
				log.Println(err)
				continue
			}
		}

		time.Sleep(1 * time.Hour)
	}
}
