package domain

type Pagination struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	ID     int    `json:"id"`
	SortBy string `json:"sort_by"`
	Search string `json:"search"`
}

type PetTypeBreedPagination struct {
	MinWeight int    `json:"min_weight"`
	MaxWeight int    `json:"max_weight"`
	Type      string `json:"type"`
	Breed     string `json:"breed"`
}
