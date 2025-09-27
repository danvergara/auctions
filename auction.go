package auctions

import (
	"errors"
	"fmt"
	"sort"
)

// Bidder struct represents the buyer willing to pay the product.
// The Name, InitialBid, MaxBid, AutoIncrement are meant to be public and edited by the user.
// The remaning maxPossibleBid and entryIndex values are meant to be manipulated by the system only.
type Bidder struct {
	Name           string
	InitialBid     int64
	MaxBid         int64
	AutoIncrement  int64
	maxPossibleBid int64
	entryIndex     int
}

// Auction struct represents the curretn Auction with a slice of bidders.
type Auction struct {
	bidders []Bidder
}

// Winner struct is the one who made it to highest Max Bid possible given the rules.
type Winner struct {
	Name   string
	MaxBid int64
}

// New function returns a pointer to an Auction, given a list of an arbitray numbers of bidders as paramenter.
func New(bidders ...Bidder) *Auction {
	return &Auction{bidders: bidders}
}

// RegisterBidder method allows the user to register a bidder individually.
func (a *Auction) RegisterBidder(b Bidder) {
	a.bidders = append(a.bidders, b)
}

// Run method runs an auction to determine a winner and the final price.
// It first validates all bidders' inputs. It then calculates each bidder's maximum possible bid
// and sorts them to find the winner. Ties are resolved by prioritizing the bidder who was registered first.
// The final winning price is calculated as the lowest possible bid the winner can make that is greater than the
// second-highest bidder's maximum possible bid.
// The time complexity is O(n log n) due to sorting the bidders.
func (a *Auction) Run() (*Winner, error) {
	if len(a.bidders) == 0 {
		return nil, errors.New("empty slice of bidders")
	}

	if err := a.validateBidders(); err != nil {
		return nil, err
	}

	// Iterate through the bidders calculate the Max Bid Possible and Auto-Increment Count to get there.
	for i := range a.bidders {
		a.bidders[i].maxPossibleBid = calculateMaxBidPossible(
			a.bidders[i].InitialBid,
			a.bidders[i].MaxBid,
			a.bidders[i].AutoIncrement,
		)

		// Keep track of the original entry order to resolve ties.
		a.bidders[i].entryIndex = i
	}

	// If there's only one bidder, the bidder should win with the InitialBid.
	if len(a.bidders) == 1 {
		return &Winner{
			Name:   a.bidders[0].Name,
			MaxBid: a.bidders[0].InitialBid,
		}, nil
	}

	// Sort the bidders based on the Max Possible Bid.
	// If the Max Possible Bids are equal, compare the original entry indexes to preserve the orginal order when it comes to ties.
	sort.Slice(a.bidders, func(i, j int) bool {
		if a.bidders[i].maxPossibleBid != a.bidders[j].maxPossibleBid {
			return a.bidders[i].maxPossibleBid > a.bidders[j].maxPossibleBid
		}

		// The earlier entry wins.
		return a.bidders[i].entryIndex < a.bidders[j].entryIndex
	})

	// Get the winner and the secod one to try to calculate the lowest amount possible to win the auction.
	winner := a.bidders[0]
	second := a.bidders[1]

	// The winner is the bidder with the lowest Bid possible to win the auction given the rules.
	w := Winner{
		Name:   a.bidders[0].Name,
		MaxBid: calculateWinnerBid(winner, second),
	}

	return &w, nil
}

// validateBidders function is used to validate the bidders data and spot nonsense.
func (a *Auction) validateBidders() error {
	for _, b := range a.bidders {
		if b.AutoIncrement <= 0 {
			return fmt.Errorf(
				"bidder '%s' has an invalid auto-increment of %d: must be a positive value",
				b.Name,
				b.AutoIncrement,
			)
		}
		if b.InitialBid < 0 {
			return fmt.Errorf(
				"bidder '%s' has a negative initial bid: %d",
				b.Name,
				b.InitialBid,
			)
		}
		if b.MaxBid < b.InitialBid {
			return fmt.Errorf(
				"bidder '%s' has a max bid (%d) lower than their initial bid (%d)",
				b.Name,
				b.MaxBid,
				b.InitialBid,
			)
		}
	}
	return nil
}

// calculateMaxBidPossible function calculates the count of auto-increment amounts the initial bid needs to get to the max bid possible,
// without esceeding the Max Bid the bidders is willing to pay.
func calculateMaxBidPossible(initialBid, maxBid, autoIncrement int64) int64 {
	// autoIncrementCount is calculated by substracting Initial Bid from the Max Bid to get the difference and divide this value by the Auto-increment amount.
	// This value means the number of auto increments to get to the Max Bid Possible without exceeding the Max Bid.
	// This is safe because division between positive integers behaves as floor division.
	autoIncrementCount := (maxBid - initialBid) / autoIncrement

	// calculateMaxBidPossible means the Maximum Bid amount possible given the rules.
	maxPossibleBid := (autoIncrementCount * autoIncrement) + initialBid

	return maxPossibleBid
}

// calculateWinnerBid calculates the winner price.
// It's the lowest amount possible to win the auction given the rules.
func calculateWinnerBid(winner, second Bidder) int64 {
	// If the winner's initial bid is already higher than the second place max, they win at their initial bid.
	if winner.InitialBid > second.maxPossibleBid {
		return winner.InitialBid
	}

	// If the Max Possible Bids are equal, just use the winners's Max Possible Bid.
	if winner.maxPossibleBid == second.maxPossibleBid {
		return winner.maxPossibleBid
	}

	// Otherwise, use the second one's Max Possible Bid to calculate the lowest possible amount to win the auction.
	// Increment one, since the winning price has to be greater than the second one's Max Possible Bid.
	incrementsNeeded := ((second.maxPossibleBid - winner.InitialBid) / winner.AutoIncrement) + 1
	return winner.InitialBid + (incrementsNeeded * winner.AutoIncrement)
}
