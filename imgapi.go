package imgapi

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"runtime"
	"encoding/json"
	"net/url"
)

// Version is the wrapper's version.
const Version = "1.0.0"

type Pong struct {
	Message string `json:"message"`
}

type Stats struct {
	Version string `json:"version"`
	Uptime int `json:"uptime"`
	Stats *runtime.MemStats `json:"stats"`
}

type Options struct {
	Port int
	Host string
	Password string
	Client *http.Client
}

type Client struct {
	Port int
	Host string
	Password string
	Client *http.Client
}


func NewClient(options *Options) *Client {
	client := &Client{}

	if options.Password != "" {
		client.Password = options.Password
	}

	if options.Host != "" {
		client.Host = options.Host
	} else {
		client.Host = "localhost"
	}

	if options.Port != 0 {
		client.Port = options.Port
	} else {
		client.Port = 3030
	}

	if options.Client != nil {
		client.Client = options.Client
	} else {
		client.Client = http.DefaultClient
	}

	return client
}

var DefaultClient = NewClient(&Options{})

func (c *Client) get(endpoint string, query url.Values) ([]byte, error) {

	if len(query) > 0 {
		endpoint = endpoint + "?" + query.Encode()
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d%s", c.Host, c.Port, endpoint), nil)

	if err != nil {
		return nil, err
	}

	if c.Password != "" {
		req.Header.Set("Authorization", c.Password)
	}

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) Ping() (*Pong, error) {
	res, err := c.get("/ping", url.Values{})

	if err != nil {
		return nil, err
	}

	pong := &Pong{}

	err = json.Unmarshal(res, pong)

	if err != nil {
		return nil, err
	}

	return pong, nil
}

func (c *Client) Stats(memStats bool) (*Stats, error) {
	value := "true"

	if memStats {
		value = "false"
	}

	res, err := c.get("/stats", url.Values{
		"noStats": {value},
	})

	if err != nil {
		return nil, err
	}

	stats := &Stats{}

	err = json.Unmarshal(res, stats)

	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (c *Client) Religion(avatar string) ([]byte, error) {
	return c.get("/religion", url.Values{
		"avatar": {avatar},
	})
}

func (c *Client) Beautiful(avatar string) ([]byte, error) {
	return c.get("/beautiful", url.Values{
		"avatar": {avatar},
	})
}

func (c *Client) Fear(avatar string) ([]byte, error) {
	return c.get("/fear", url.Values{
		"avatar": {avatar},
	})
}

func (c *Client) Sacred(avatar string) ([]byte, error) {
	return c.get("/sacred", url.Values{
		"avatar": {avatar},
	})
}

func (c *Client) Painting(avatar string) ([]byte, error) {
	return c.get("/painting", url.Values{
		"avatar": {avatar},
	})
}

// Color is either a hex like '#FFFFFF' (# is optional) or a name like 'blue'
func (c *Client) Color(color string) ([]byte, error) {
	return c.get("/color", url.Values{
		"color": {color},
	})
}

func (c *Client) Delete(avatar string) ([]byte, error) {
	return c.get("/delete", url.Values{
		"avatar": {avatar},
	})
}

func (c *Client) Garbage(avatar string) ([]byte, error) {
	return c.get("/garbage", url.Values{
		"avatar": {avatar},
	})
}

func (c *Client) Tom(avatar string) ([]byte, error) {
	return c.get("/tom", url.Values{
		"avatar": {avatar},
	})
}

func (c *Client) Bed(avatar, target string) ([]byte, error) {
	return c.get("/bed", url.Values{
		"avatar": {avatar},
		"target": {target},
	})
}

func (c *Client) Crush(avatar, target string) ([]byte, error) {
	return c.get("/crush", url.Values{
		"avatar": {avatar},
		"target": {target},
	})
}

func (c *Client) Dipshit(text string) ([]byte, error) {
	return c.get("/dipshit", url.Values{
		"text": {text},
	})
}

func (c *Client) Picture(avatar string) ([]byte, error) {
	return c.get("/picture", url.Values{
		"avatar": {avatar},
	})
}

// Max char limit 165
func (c *Client) Tweet(text string) ([]byte, error) {
	return c.get("/tweet", url.Values{
		"text": {text},
	})
}

func (c *Client) Truth(avatar string) ([]byte, error) {
	return c.get("/truth", url.Values{
		"avatar": {avatar},
	})
}

func (c *Client) Mask(avatar string) ([]byte, error) {
	return c.get("/mask", url.Values{
		"avatar": {avatar},
	})
}

// Max char limit 41
func (c *Client) Father(avatar, text string) ([]byte, error) {
	return c.get("/father", url.Values{
		"avatar": {avatar},
		"text": {text},
	})
}

// Max char limit 21
func (c *Client) Archievement(avatar, text string) ([]byte, error) {
	return c.get("/achievement", url.Values{
		"avatar": {avatar},
		"text": {text},
	})
}
