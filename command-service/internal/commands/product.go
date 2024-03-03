package commands

type CreateProduct struct {
	BucketId    string
	Name        string
	Description string
	Price       float32
	Quantity    int32
}

type CreateProductOut struct {
	ProductId string `json:"productId"`
}

type UpdateProduct struct {
	BucketId    string
	ProductId   string
	Name        string
	Description string
	Price       float32
}

type DeleteProduct struct {
	BucketId  string
	ProductId string
}
