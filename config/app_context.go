package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type AppContext struct {
	EnvVars Environment
}

func NewAppContext() AppContext {
	var environment Environment
	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}

	err = env.Unmarshal(es, &environment)
	if err != nil {
		log.Fatal(err)
	}

	appCtx := AppContext{
		EnvVars: environment,
	}

	return appCtx
}
