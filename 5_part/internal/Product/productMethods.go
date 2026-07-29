package product

type Product struct {
	Name  string
	Price float32
}

var products = []Product{
	{
		Name:  "Молоко",
		Price: 110,
	},
	{
		Name:  "Хлеб",
		Price: 23.15,
	},
	{
		Name:  "Макароны",
		Price: 45.26,
	},
	{
		Name:  "Соль",
		Price: 17.34,
	},
	{
		Name:  "Сахарок",
		Price: 67.69,
	},
}

func GetProductList() []Product {
	return products
}
