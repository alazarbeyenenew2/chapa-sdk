package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	DoGet(url string, headers map[string]string, response interface{}) error
	DoPost(url string, request interface{}, headers map[string]string, response interface{}) error
}

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) DoGet(url string, headers map[string]string, response interface{}) error {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create GET request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GET request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read GET response body: %w", err)
	}

	if err := json.Unmarshal(body, response); err != nil {
		return fmt.Errorf("failed to unmarshal GET response: %w", err)
	}

	return nil
}

func (c *Client) DoPost(url string, request interface{}, headers map[string]string, response interface{}) error {

	body, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal POST request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create POST request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("POST request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read POST response body: %w", err)
	}

	if err := json.Unmarshal(body, response); err != nil {
		return fmt.Errorf("failed to unmarshal POST response: %w", err)
	}

	return nil
}
