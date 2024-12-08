package json_placeholder_client

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type JSONPlaceholderClient struct {
	rwMu           sync.RWMutex
	ResultsStorage []string
}

func (c *JSONPlaceholderClient) validateId(idx int) error {
	if idx < 0 || idx >= len(c.ResultsStorage) {
		return fmt.Errorf("index out of range: %d", idx)
	}
	return nil
}

func (c *JSONPlaceholderClient) GetResultById(idx int) (string, error) {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()

	if err := c.validateId(idx); err != nil {
		return "", err
	}
	return c.ResultsStorage[idx], nil
}

func (c *JSONPlaceholderClient) Add(idx int, payload string) error {
	c.rwMu.RLock()
	if err := c.validateId(idx); err != nil {
		return err
	}
	c.rwMu.RUnlock()

	c.ResultsStorage[idx] = payload
	return nil
}

func (c *JSONPlaceholderClient) AddFromUrl(idx int, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error during GET request: %s", err)
	}

	payload, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("error during reading response body: %s", err)
	}

	c.rwMu.RLock()
	if err := c.validateId(idx); err != nil {
		return err
	}
	c.rwMu.RUnlock()

	c.ResultsStorage[idx] = string(payload)
	return nil
}

func NewJSONPlaceholderClient(resultsStorageSize int) *JSONPlaceholderClient {
	return &JSONPlaceholderClient{
		ResultsStorage: make([]string, resultsStorageSize),
	}
}
