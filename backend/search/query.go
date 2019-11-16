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


func (q *Query) Search() []blockchain.Block{
	blocks := make([]blockchain.Block,0)
	blocks = append(blocks, blockchain.Block{
		FirstName:   "Bobby",
		Name:        "Bob",
		Location:    "EPFL",
		Success:     true,
		ReportHash:  "B94D27B9934D3E08A52E52D7DA7DABFAC484EFE37A5380EE9088F7ACE2EFCDE9",
		Total:       1,
		Succeed:     1,
		SuccessRate: 1.0,
	})
	//TODO: implements real search on blockchain
	return blocks
}