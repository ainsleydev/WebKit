package wordpress

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ainsleydev/webkit/pkg/util/httputil"
)

// Client is a WordPress API client.
type Client struct {
	opts   Options
	client *http.Client
	reader func(io.Reader) ([]byte, error)
}

// New creates a new WordPress API client.
func New(opts Options) (*Client, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}
	return &Client{
		opts:   opts,
		client: &http.Client{},
		reader: io.ReadAll,
	}, err
}

// Collection names for WordPress that are used in the API
// after the base URL.
const (
	CollectionUsers      = "users"
	CollectionPosts      = "posts"
	CollectionPages      = "pages"
	CollectionMedia      = "media"
	CollectionMeta       = "meta"
	CollectionRevisions  = "revisions"
	CollectionComments   = "comments"
	CollectionTaxonomies = "taxonomies"
	CollectionTerms      = "terms"
	CollectionStatuses   = "statuses"
	CollectionTypes      = "types"
)

// Get sends a GET request to the specified WordPress URL and returns the response body.
// The base URL is prepended to the URL, for example:
// https://wordpress/wp-json/wp/v2/posts/21
func (c *Client) Get(ctx context.Context, url string) ([]byte, error) {
	path := fmt.Sprintf("%s/%s", c.opts.baseURL, strings.TrimLeft(url, "/"))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	if c.opts.hasBasicAuth {
		req.SetBasicAuth(c.opts.authUser, c.opts.authPassword)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := c.reader(resp.Body)
	if err != nil {
		return nil, err
	}

	if !httputil.Is2xx(resp.StatusCode) {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return body, nil
}

// GetAndUnmarshal performs an HTTP GET request to the specified WordPress URL,
// unmarshal the response body into the provided struct type, and returns any error.
func (c *Client) GetAndUnmarshal(ctx context.Context, url string, v any) error {
	body, err := c.Get(ctx, url)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return err
	}
	return nil
}
