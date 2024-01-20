package bot

import (
	"fmt"

	"github.com/google/go-querystring/value"
)

type ShopeeClient struct {
	affiliateID string
}

func NewShopeeClient(config *Config) *ShopeeClient {
	return &ShopeeClient{affiliateID: config.Shopee.AffiliateID}
}

func (c *ShopeeClient) CreateAffiliateURL(productID string) string {
	v := value.Values{
		"affiliate_id": c.affiliateID,
		"itemid":       productID,
	}

	return fmt.Sprintf("https://shopee.com.my/%s?%s", productID, v.Encode())
}
