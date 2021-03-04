package main

import (
	"context"
	"fmt"
	"os"

	"github.com/animeshon/gqlgen/api"
	"github.com/animeshon/gqlgenc/clientgen"
	"github.com/animeshon/gqlgenc/clientgenv2"
	"github.com/animeshon/gqlgenc/config"
	"github.com/animeshon/gqlgenc/generator"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err.Error())
		os.Exit(2)
	}

	clientGen := api.AddPlugin(clientgen.New(cfg.Query, cfg.Client, cfg.Generate))
	if cfg.Generate != nil {
		if cfg.Generate.ClientV2 {
			clientGen = api.AddPlugin(clientgenv2.New(cfg.Query, cfg.Client, cfg.Generate))
		}
	}

	if err := generator.Generate(ctx, cfg, clientGen); err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err.Error())
		os.Exit(4)
	}
}
