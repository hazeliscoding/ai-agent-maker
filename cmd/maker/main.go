package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hazeliscoding/ai-agent-maker/internal/agents"
)

func main() {
	key := os.Getenv("OPENAI_API_KEY")
	fmt.Println("KEY FROM ENV:", key)

	ctx := context.Background()
	prompt := "Write a simple todo program in Go."
	apiClient := agents.NewOpenAI(ctx, key, nil)

	res, err := apiClient.Query("", prompt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%+v\n\n", res)
}
