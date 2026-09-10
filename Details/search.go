package details

import {
	"fmt"
	"time"
	"log"
	"net/http"
}

type response struct {
	title		string 		`json:"title"`
	location	string 		`json:"location"`
	date		time.Time	`json:"date"`
	seat[2][2]  int			`json:"seat"`
	avail		bool		`json:"avail"`
	details     string      `json:"detail"`
	id			int			`json:"id"`
}

type request struct {
	title		string		`json:title`
}

func getinfo(ctx context.Context){

}

func eventinfo(){

}

func main(){
	mux := http.NewServeMux

	mux.HandleFunc("/search", getinfo)
	mux.HandleFunc("/event/{id}", eventinfo)

	fmt.Printf("Server has started")
	log.Fatal(http.ListenAndServe(":8080",mux))

}




