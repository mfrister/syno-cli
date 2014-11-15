package synoapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	api_base string
	sid      string
}

func NewClient(base_url string) *Client {
	if !strings.HasPrefix(base_url, "https://") {
		log.Printf("synoapi.Client: Warning: Base URL does not use HTTPS: '%s'", base_url)
	}
	api_base := fmt.Sprintf("%s/webapi", base_url)

	return &Client{api_base: api_base}
}

type SynoLoginResponse struct {
	synoBaseResponse
	Data struct {
		Sid string
	}
}

func (c *Client) Login(user string, password string) error {
	params := map[string]string{
		"account": user,
		"passwd":  password,
		"format":  "sid",
	}

	resp := SynoLoginResponse{synoBaseResponse: synoBaseResponse{}}
	err := c.request("auth.cgi", "SYNO.API.Auth", "3", "login", params, &resp)

	if err != nil {
		return err
	}

	log.Printf("synoapi.Client.Login: Login successful with sid '%s'", resp.Data.Sid)
	c.sid = resp.Data.Sid

	return err
}

func (c *Client) request(path string, api string, api_version string, method string, params map[string]string, r SynoBaseResponse) ClientError {
	p := url.Values{}
	p.Set("api", api)
	p.Set("version", api_version)
	p.Set("method", method)
	for key, value := range params {
		p.Set(key, value)
	}
	if c.sid != "" {
		p.Set("_sid", c.sid)
	}
	url := fmt.Sprintf("%s/%s?%s", c.api_base, path, p.Encode())

	log.Printf("synoapi.Client: GET %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return NewClientError("HTTP request failed", err)
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&r)
	if err != nil {
		return NewClientError("JSON decoding failed", err)
	}
	log.Printf("synoapi.Client: Response: %v code: %v", r, resp.StatusCode)
	if !r.Successful() {
		return NewClientError(fmt.Sprintf("Synology API returned error code %v", r.ErrorCode()), nil)
	}
	return nil
}
