package aggs

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	getAggsPath           = "/v2/aggs/ticker/{ticker}/range/{multiplier}/{resolution}/{from}/{to}"
	getPreviousClosePath  = "/v2/aggs/ticker/{ticker}/prev"
	getGroupedDailyPath   = "/v2/aggs/grouped/locale/{locale}/market/{marketType}/{date}"
	getDailyOpenClosePath = "/v1/open-close/{ticker}/{date}"
)

// Client defines a REST client for the Polygon aggregates API.
type Client struct {
	client.Client
}

// GetAggs retrieves aggregate bars for a specified ticker over a given date range in custom time window sizes.
// For example, if timespan = ‘minute’ and multiplier = ‘5’ then 5-minute bars will be returned.
func (ac *Client) GetAggs(ctx context.Context, params models.GetAggsParams, opts ...client.Option) (*models.AggsResponse, error) {
	res := &models.AggsResponse{}
	err := ac.Call(ctx, http.MethodGet, getAggsPath, params, res, opts...)
	return res, err
}

// GetPreviousClose retrieves the previous day's open, high, low, and close (OHLC) for the specified ticker.
func (ac *Client) GetPreviousClose(ctx context.Context, params models.GetPreviousCloseParams, opts ...client.Option) (*models.AggsResponse, error) {
	res := &models.AggsResponse{}
	err := ac.Call(ctx, http.MethodGet, getPreviousClosePath, params, res, opts...)
	return res, err
}

// GetGroupedDaily retrieves the daily open, high, low, and close (OHLC) for the specified market type.
func (ac *Client) GetGroupedDaily(ctx context.Context, params models.GetGroupedDailyParams, opts ...client.Option) (*models.AggsResponse, error) {
	res := &models.AggsResponse{}
	err := ac.Call(ctx, http.MethodGet, getGroupedDailyPath, params, res, opts...)
	return res, err
}

// GetDailyOpenClose retrieves the open, close and afterhours prices of a specific symbol on a certain date.
func (ac *Client) GetDailyOpenClose(ctx context.Context, params models.GetDailyOpenCloseParams, opts ...client.Option) (*models.DailyOpenCloseResponse, error) {
	res := &models.DailyOpenCloseResponse{}
	err := ac.Call(ctx, http.MethodGet, getDailyOpenClosePath, params, res, opts...)
	return res, err
}
