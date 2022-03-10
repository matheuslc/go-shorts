package shorturl

import (
	"fmt"
	"net/http"

	"github.com/matheuslc/go-shorts/config"
)

func Handler(w http.ResponseWriter, r *http.Request, ctx config.AppContext) {
	fmt.Fprintf(w, "Hi there, I love %s!", ctx.EnvVars.Redis.RedisAddress)
}
