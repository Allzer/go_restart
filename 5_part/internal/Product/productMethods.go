package product

type Product struct {
	Name  string
	Price float32
}

func GetProductList() []Product {
	var productList []Product
	productList = append(productList,
		Product{
			Name:  "Молоко",
			Price: 110.0,
		},
		Product{
			Name:  "Хлеб",
			Price: 23.15,
		},
		Product{
			Name:  "Макароны",
			Price: 45.26,
		},
		Product{
			Name:  "Соль",
			Price: 17.34,
		},
		Product{
			Name:  "Сахарок",
			Price: 67.69,
		},
	)
	return productList
}
