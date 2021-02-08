package trading212

import (
	"context"
	"net/http"
	"time"
)

const (
	OrderTypeMarket    OrderType = "MARKET"
	OrderTypeStop      OrderType = "STOP"
	OrderTypeLimit     OrderType = "LIMIT"
	OrderTypeStopLimit OrderType = "STOP_LIMIT"

	OrderTimeValidityDay           OrderTimeValidity = "DAY"
	OrderTimeValidityTillCancelled OrderTimeValidity = "GOOD_TILL_CANCEL"
)

// OrderType is the type of order e.g OrderTypeStop/OrderTypeMarket.
type OrderType string

// OrderTimeValidity is how long the order is valid for e.g
// OrderTimeValidityDay/OrderTimeValidityTilCancelled.
type OrderTimeValidity string

// CreateOrderRequest contains the request parameters to create an order.
type CreateOrderRequest struct {
	InstrumentCode string            `json:"instrumentCode"`
	OrderType      OrderType         `json:"orderType"`
	StopPrice      float64           `json:"stopPrice,omitempty"`
	LimitPrice     float64           `json:"limitPrice,omitempty"`
	Quantity       float64           `json:"quantity"`
	TimeValidity   OrderTimeValidity `json:"timeValidity"`
}

// AccountResponse returned from various v2 endpoints.
type AccountResponse struct {
	Account struct {
		Dealer    string `json:"dealer"`
		Positions []struct {
			PositionID            string    `json:"positionId"`
			HumanID               string    `json:"humanId"`
			Created               time.Time `json:"created"`
			AveragePrice          float64   `json:"averagePrice"`
			AveragePriceConverted float64   `json:"averagePriceConverted"`
			CurrentPrice          float64   `json:"currentPrice"`
			Value                 float64   `json:"value"`
			Investment            float64   `json:"investment"`
			Code                  string    `json:"code"`
			Margin                float64   `json:"margin"`
			Ppl                   float64   `json:"ppl"`
			Quantity              float64   `json:"quantity"`
			MaxBuy                float64   `json:"maxBuy"`
			MaxSell               float64   `json:"maxSell"`
			MaxOpenBuy            float64   `json:"maxOpenBuy"`
			MaxOpenSell           float64   `json:"maxOpenSell"`
			Frontend              string    `json:"frontend"`
			AutoInvestQuantity    float64   `json:"autoInvestQuantity"`
			FxPpl                 float64   `json:"fxPpl"`
		} `json:"positions"`
		Cash struct {
			Free                 float64 `json:"free"`
			Total                float64 `json:"total"`
			Interest             float64 `json:"interest"`
			Indicator            float64 `json:"indicator"`
			Commission           float64 `json:"commission"`
			Cash                 float64 `json:"cash"`
			Ppl                  float64 `json:"ppl"`
			Result               float64 `json:"result"`
			SpreadBack           float64 `json:"spreadBack"`
			NonRefundable        float64 `json:"nonRefundable"`
			Dividend             float64 `json:"dividend"`
			StockInvestment      float64 `json:"stockInvestment"`
			FreeForStocks        float64 `json:"freeForStocks"`
			TotalCashForWithdraw float64 `json:"totalCashForWithdraw"`
			BlockedForStocks     float64 `json:"blockedForStocks"`
			PieCash              float64 `json:"pieCash"`
		} `json:"cash"`
		LimitStop    []interface{} `json:"limitStop"`
		Oco          []interface{} `json:"oco"`
		IfThen       []interface{} `json:"ifThen"`
		EquityOrders []struct {
			OrderID        string    `json:"orderId"`
			Type           string    `json:"type"`
			Code           string    `json:"code"`
			Quantity       float64   `json:"quantity"`
			FilledQuantity float64   `json:"filledQuantity"`
			Status         string    `json:"status"`
			StopPrice      float64   `json:"stopPrice"`
			Created        time.Time `json:"created"`
			Frontend       string    `json:"frontend"`
		} `json:"equityOrders"`
		EquityValueOrders []interface{} `json:"equityValueOrders"`
		ID                int           `json:"id"`
		Timestamp         int64         `json:"timestamp"`
	} `json:"account"`
}

// CreateOrder creates a new order.
func (c *Client) CreateOrder(
	ctx context.Context, createReq CreateOrderRequest,
) (AccountResponse, error) {
	endpoint := "/rest/public/v2/equity/order"
	req, err := c.newRequest(ctx, http.MethodPost, endpoint, createReq)
	if err != nil {
		return AccountResponse{}, err
	}

	var resp AccountResponse
	return resp, c.do(req, &resp)
}

// UpdateOrderRequest contains the request parameters to update an order.
type UpdateOrderRequest struct {
	StopPrice    float64           `json:"stopPrice,omitempty"`
	LimitPrice   float64           `json:"limitPrice,omitempty"`
	Quantity     int               `json:"quantity,omitempty"`
	TimeValidity OrderTimeValidity `json:"timeValidity,omitempty"`
}

// UpdateOrder with id.
func (c *Client) UpdateOrder(
	ctx context.Context, id string, updateReq UpdateOrderRequest,
) (AccountResponse, error) {
	endpoint := "/rest/public/v2/equity/order/" + id
	req, err := c.newRequest(ctx, http.MethodPut, endpoint, updateReq)
	if err != nil {
		return AccountResponse{}, err
	}

	var resp AccountResponse
	return resp, c.do(req, &resp)
}

// CancelOrder with id.
func (c *Client) CancelOrder(
	ctx context.Context, id string,
) (AccountResponse, error) {
	endpoint := "/rest/public/v2/equity/order/" + id
	req, err := c.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return AccountResponse{}, err
	}

	var resp AccountResponse
	return resp, c.do(req, &resp)
}
