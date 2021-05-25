package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"context"
	"fmt"
	"myapp/rpc/dummy"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//Open Grpc Client
	dumCon, con := dummy.Connect()
	defer con.Close()

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {

		//Get Parameter
		params := r.URL.Query()
		fmt.Println("Parameter in Query : ", params)

		a := params.Get("a")
		b := params.Get("b")

		//Convert to Int
		c, _ := strconv.Atoi(a)
		d, _ := strconv.Atoi(b)

		//grpc Connect
		resp, err := dumCon.DummyRpc(context.Background(), &dummy.DummyRpcParam{A: int32(c), B: int32(d)})
		if err != nil {
			panic(err)
		}

		//Return the GRPC Result
		fmt.Fprint(w, resp.Result)
	})

	log.Println("Listen and Serve at http://localhost:" + port)
	//Listen and Serve http
	http.ListenAndServe(":"+port, nil)
}
