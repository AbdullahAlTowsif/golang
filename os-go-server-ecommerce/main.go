// go mod init moduleName
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm towsif. I'm a student. Want to become a software engineer.")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Please give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Please give me POST request", 400)
		return
	}

	/*
		1. take body information (description, imageUrl, price, title) from r.Body
		2. create an instance using Product struct with the body information
		3. append the instance into productList
	*/

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	// frontend theke json asbe oitake amra decode kore struct/js object e convert krbo.
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	w.WriteHeader(201)

	// backend theke kono information jodi JSON e pathaite chai tahole encoder
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}

func main() {
	mux := http.NewServeMux() //router --> mux

	// route
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-product", createProduct)

	fmt.Println("Server running on: 5000")

	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
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

/*
How it works:
* Client request korse oita gese router er kache.
* Router GO er j server acher oitar bhitor j Network Interface Card (NIC) ache oitar kache request
forward korse. NIC electro magnetic wave k binary te convert krlo. Eibar oi data ta niye rakhlo
RAM e. RAM er j NIC WRITE BUFFER ache tar kache data(binary) ta niye rakhlo.
* Ekhon NIC kernel k interrupt korbe. Kernel k interrupt kora mane holo kernel bujhe gelo tar kache
ekta request ashce. Tokhon kernel oi data ta read korbe.
* Kernel data ta read kore dekhlo kon PORT e data ta ashce. Oi PORT RAM er arekta buffer Socker
er jonno allocated.
* Tarpor kernel oi socket er jonno j allocate kora RECIEVE BUFFER ache sekhane data ta rakhbe.
Data rakhar shate shate FILE DESCRIPTOR e jabe. File Descriptor e giye Oi socket er jonno j file
descriptor ache suppose 8 oitak mark kore dbe. Tokhon oita readble hobe. Readbale Information ta jabe
GO RUNTIME er kache
* GO RUNTIME dekhe j kon GO ROUTIME SOCKET(MAIN) ta banaite bolsilo. Tokhon MAIN GO ROUTINE k bole j
8 READABLE.
* Ekhon MAIN GO ROUTINE accept function k call dey and direct kernel k bole j he want to read 8.
* Tarpor socket er j RECEIVE BUFFER ache shetar information ta MAIN GO ROUTINE read kore. Data ta
pailo accept er bhitor theke.
* Tarpor MAIN GO ROUTINE bole GO RUNTIME k ami j data ta paisi sheta handle korar jonno arekta
GO ROUTINE create kore dao.
* Oi new created GO ROUTINE khuje dekhe mux/route means kon route er jonno request ta ashcilo.
Tarpor oi route er jonno corresponding funtion ta execute kore. Tarpor info write kore.
Oi socket er arekta buffer ache SEND BUFFER jkhane data ta write kore dey.
* Tarpor socket kernel k interrupt kore.
* Tarpor kernel RAM er moddhe RING BUFFER ache oikhane data ta write kore.
* Finally NIC RING BUFFER theke data ta read kore niye electromagnetic wave akare ROUTER er kache pathai.
* And finally router client er kache data ta pathai.

*/

/*
GO RUNTIME



*/
