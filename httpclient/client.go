package httpclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AccessToken string
}

func NewClientWithToken(token string) *Client {
	return &Client{
		AccessToken: token,
	}
}

func (c *Client) Get(url string, output interface{}) error {
	httpClient := &http.Client{}
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return e
	}
	if c.AccessToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	}

	resp, e := httpClient.Do(req)
	if e != nil {
		return e
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		e = json.Unmarshal(bodyBytes, output)
		if e != nil {
			return e
		}
		return nil
	}
	return fmt.Errorf("error: %s", resp.Status)
}
