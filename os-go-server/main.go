// go mod init moduleName
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm towsif. I'm a student. Want to become a software engineer.")
}

func main() {
	mux := http.NewServeMux() //router --> mux
	
	// route
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on: 3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
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