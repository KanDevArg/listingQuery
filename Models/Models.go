package Models

type Hit struct {
	Location string	`json:"location"`
}

type ListingSearch struct {
	Location string	`json:"title"`
}

func (l *ListingSearch) TableName() string {
	return "listingSearch"
}

func GetAllHits(hits *[]Hit) (err error) {
	hits = &[]Hit{}
	return nil
}
