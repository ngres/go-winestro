package winestro

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const host = "https://weinstore.net/xml/v22.0/wbo-API.php"

type Config struct {
	UID    int    `mapstructure:"uid" structs:"uid" json:"uid" jsonschema:"title=User ID,description=Identifies the Winestro tenant.,example=1000" env:"WBO_UID"`
	User   string `mapstructure:"api_user" structs:"api_user" json:"user" jsonschema:"title=API user,description=The API user name.,example=apiuser-1000" env:"WBO_API_USER"`
	Code   string `mapstructure:"api_code" structs:"api_code" json:"code" jsonschema:"title=API code,description=The API user secret." env:"WBO_API_CODE"`
	ShopID int    `mapstructure:"shop_id" structs:"shop_id" json:"shop_id" jsonschema:"title=Shop ID,description=Identifies the shop.,example=1" env:"WBO_SHOP_ID"`
}

type Client struct {
	httpClient *http.Client
	config     *Config
}

func NewClientFromConfig(config *Config) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		config: config,
	}
}

func NewClient(uid int, user string, code string, shopID int) *Client {
	return NewClientFromConfig(&Config{
		UID:    uid,
		User:   user,
		Code:   code,
		ShopID: shopID,
	})

}

func NewClientFromEnv() *Client {
	uid, err := strconv.Atoi(os.Getenv("WBO_UID"))
	if err != nil {
		return nil
	}
	shopID, err := strconv.Atoi(os.Getenv("WBO_SHOP_ID"))
	if err != nil {
		return nil
	}
	return NewClient(uid, os.Getenv("WBO_API_USER"), os.Getenv("WBO_API_CODE"), shopID)
}

func (c *Client) do(action string, params map[string]string, v interface{}) error {
	req, err := http.NewRequest("POST", host, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("UID", fmt.Sprint(c.config.UID))
	q.Add("apiUSER", c.config.User)
	q.Add("apiCODE", c.config.Code)
	q.Add("apiShopID", fmt.Sprint(c.config.ShopID))
	q.Add("apiACTION", action)
	q.Add("output", "xml")

	for key, value := range params {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return nil
	}
	return decode(resp.Body, v)
}

func decode(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}
