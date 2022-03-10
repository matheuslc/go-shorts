package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matheuslc/go-shorts/config"
)

func handler(w http.ResponseWriter, r *http.Request, ctx config.AppContext) {
	fmt.Fprintf(w, "Hi there, I love %s!", ctx.EnvVars.Redis.RedisAddress)
}

func main() {
	appContext := config.NewAppContext()

	fmt.Println(appContext.EnvVars.Redis.RedisAddress)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, appContext)
	})
	log.Fatal(http.ListenAndServe(":3000", nil))

}
