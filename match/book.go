package match

import (
	"github.com/mit-dci/lit/crypto/koblitz"
)

// LimitOrderbook is the interface for a limit order book.
// The difference between the order book and the matching engine is that the matching engine is what processes orders and generates executions, whereas the limit orderbook does not.
// The order book takes in executions, and allows read access to the state of the orderbook.
// This can only be updated using the outputs of the matching engine.
type LimitOrderbook interface {
	// UpdateBookExec takes in an order execution and updates the orderbook.
	UpdateBookExec(orderExec *OrderExecution) (err error)
	// UpdateBookCancel takes in an order cancellation and updates the orderbook.
	UpdateBookCancel(cancel *CancelledOrder) (err error)
	// UpdateBookPlace takes in an order, ID, timestamp, and adds the order to the orderbook.
	UpdateBookPlace(limitIDPair *LimitOrderIDPair) (err error)
	// GetOrder gets an order from an OrderID
	GetOrder(orderID *OrderID) (limOrder *LimitOrderIDPair, err error)
	// CalculatePrice takes in a pair and returns the calculated price based on the orderbook.
	CalculatePrice() (price float64, err error)
	// GetOrdersForPubkey gets orders for a specific pubkey.
	GetOrdersForPubkey(pubkey *koblitz.PublicKey) (orders map[float64][]*LimitOrderIDPair, err error)
	// ViewLimitOrderbook takes in a trading pair and returns the orderbook as a map
	ViewLimitOrderBook() (book map[float64][]*LimitOrderIDPair, err error)
}

// AuctionOrderbook is the interface for an auction order book.
// The difference between the order book and the matching engine is that the matching engine is what processes orders and generates executions, whereas the limit orderbook does not.
// The order book takes in executions, and allows read access to the state of the orderbook.
type AuctionOrderbook interface {
	// UpdateBookExec takes in an order execution and updates the orderbook.
	UpdateBookExec(orderExec *OrderExecution) (err error)
	// UpdateBookCancel takes in an order cancellation and updates the orderbook.
	UpdateBookCancel(cancel *CancelledOrder) (err error)
	// UpdateBookPlace takes in an order, ID, auction ID and adds the order to the orderbook.
	UpdateBookPlace(auctionIDPair *AuctionOrderIDPair) (err error)
	// GetOrder gets an order from an OrderID
	GetOrder(orderID *OrderID) (limOrder *AuctionOrderIDPair, err error)
	// CalculatePrice takes in a pair and returns the calculated price based on the orderbook.
	// This only works for a specific auction
	CalculatePrice(auctionID *AuctionID) (price float64, err error)
	// GetOrdersForPubkey gets orders for a specific pubkey.
	GetOrdersForPubkey(pubkey *koblitz.PublicKey) (orders map[float64][]*AuctionOrderIDPair, err error)
	// ViewAuctionOrderBook takes in a trading pair and returns the orderbook as a map
	ViewAuctionOrderBook() (book map[float64][]*AuctionOrderIDPair, err error)
}
