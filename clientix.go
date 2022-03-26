package clientix

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const DefaultDomain = "klientiks.ru"
const DefaultTimeout = time.Second * 30

type Client struct {
	opts Opts

	httpClient  *http.Client
	rateLimiter *rate.Limiter
}

type Opts struct {
	Timeout     time.Duration
	Domain      string
	AccountId   string
	UserId      string
	AccessToken string
}

func New(opts Opts) *Client {
	return &Client{
		opts: opts,

		httpClient:  http.DefaultClient,
		rateLimiter: rate.NewLimiter(rate.Every(time.Second), 2),
	}
}

func (c *Client) GetTimeout() time.Duration {
	if c.opts.Timeout != 0 {
		return c.opts.Timeout
	}

	return DefaultTimeout
}

func (c *Client) GetDomain() string {
	if c.opts.Domain != "" {
		return c.opts.Domain
	}

	return DefaultDomain
}

func (c *Client) GetAccountId() string {
	return c.opts.AccountId
}

func (c *Client) GetUserId() string {
	return c.opts.UserId
}

func (c *Client) GetAccessToken() string {
	return c.opts.AccessToken
}

func (c *Client) HttpRequest(ctx context.Context, method, url string, body io.Reader) (data []byte, err error) {
	if err = c.rateLimiter.Wait(ctx); err != nil {
		return
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, c.opts.Timeout)
	defer cancel()

	var req *http.Request
	if req, err = http.NewRequestWithContext(timeoutCtx, method, url, body); err != nil {
		return
	}
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	var res *http.Response
	if res, err = c.httpClient.Do(req); err != nil {
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
