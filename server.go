package main

import (
	"context"
	"fmt"
	"log"
	"myapp/rpc/math"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mathCon, con := math.Connect()
		defer con.Close()

		resp, err := mathCon.Plus(context.Background(), &math.PlusParam{
			A: 5,
			B: 4,
		})

		if err != nil {
			fmt.Fprintf(w, "err: %+v", err)
		}

		fmt.Fprintf(w, "%+v", resp)
	})

	log.Println(http.ListenAndServe(":"+port, router))
}
