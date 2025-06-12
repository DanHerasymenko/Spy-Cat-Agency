package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CatBreed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CatAPI struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewCatAPI(baseURL, apiKey string) *CatAPI {
	return &CatAPI{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c *CatAPI) ValidateBreed(ctx context.Context, breed string) (bool, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/breeds", c.baseURL), nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("x-api-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var breeds []CatBreed
	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		return false, fmt.Errorf("failed to decode response: %w", err)
	}

	for _, b := range breeds {
		if b.Name == breed {
			return true, nil
		}
	}

	return false, nil
}
