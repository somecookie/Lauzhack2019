package search

import "github.com/somecookie/Lauzhack2019/backend/blockchain"
type Query struct {
	FirstName      string
	Name           string
	Location       string
	IsSuccess      bool
	MinSuccessRate float64
	ReportHash     string
}

func (q *Query) Search() []blockchain.Block {

	return nil
}
