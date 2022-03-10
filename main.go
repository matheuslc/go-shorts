package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matheuslc/go-shorts/config"
	"github.com/matheuslc/go-shorts/internal/shorturl"
)

func main() {
	appContext := config.NewAppContext()

	fmt.Println(appContext.EnvVars.Redis.RedisAddress)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shorturl.Handler(w, r, appContext)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
