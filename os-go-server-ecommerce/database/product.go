package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productId int) *Product {
	for _, product := range productList {
		if product.ID == productId {
			return &product
		}
	}

	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productId int) {
	var tempList []Product = make([]Product, 0)

	for _, product := range productList {
		if product.ID != productId {
			tempList = append(tempList, product)
		}
	}
	productList = tempList
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is Yummy! I love orange",
		Price:       100,
		ImgUrl:      "https://www.chemwatch.net/wp-content/uploads/2021/11/image-6-1024x682.jpeg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is Yummy! I love Apple",
		Price:       120,
		ImgUrl:      "https://cdn.mos.cms.futurecdn.net/38Moe9n3b72uVEf2Ti7KmV.jpg",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Yummy! I love Banana",
		Price:       20,
		ImgUrl:      "https://www.dole.com/sites/default/files/media/2025-01/banana-cavendish_0.png",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Grape",
		Description: "Grape is Yummy! I love Grape",
		Price:       180,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Grapes%2C_Rostov-on-Don%2C_Russia.jpg/1280px-Grapes%2C_Rostov-on-Don%2C_Russia.jpg",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "Mango is my favorite! I love Mango",
		Price:       80,
		ImgUrl:      "https://png.pngtree.com/png-vector/20250303/ourmid/pngtree-ripe-mango-fruit-with-leaf-for-healthy-snack-png-image_15699037.png",
	}

	prd6 := Product{
		ID:          6,
		Title:       "Jackfruit",
		Description: "Jackfruit is Yummy! I love Jackfruit",
		Price:       140,
		ImgUrl:      "https://www.gardenia.net/wp-content/uploads/2025/05/shutterstock_2453997129.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)
}
