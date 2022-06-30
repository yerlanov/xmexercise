package ipapi

import (
	"context"
	"github.com/go-resty/resty/v2"
	"time"
)

type client struct {
	client *resty.Client
}

func NewClient(baseUrl string) Service {
	s := new(client)
	s.client = resty.New().
		SetBaseURL(baseUrl).
		SetTimeout(15 * time.Second)
	return s
}

type Service interface {
	GetIPInformation(ctx context.Context, ip string) (res IpApi, err error)
}

func (c *client) GetIPInformation(ctx context.Context, ip string) (res IpApi, err error) {
	params := map[string]string{
		"ip": ip,
	}

	resp, err := c.client.R().
		SetPathParams(params).
		SetContext(ctx).
		SetResult(&res).
		Get("/{ip}/json/")
	if err != nil {
		return
	}

	if resp.Error() != nil {
		return
	}
	return
}
