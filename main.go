package main

import (
	"context"
	"fmt"
	"os"

	"github.com/moby/buildkit/client/llb"
)

func main() {
	alpine := llb.Image("docker.io/library/alpine:latest")

	builder := alpine.Run(
		llb.Shlex("echo 'Hello, BuildKit!' > /hello"),
	)

	finalState := builder.Root()

	def, err := finalState.Marshal(
		context.Background(),
		llb.LinuxAmd64,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling LLB: %v\n", err)
		os.Exit(1)
	}

	llb.WriteTo(def, os.Stdout)
}
