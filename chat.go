package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a prompt as a command-line argument.")
		return
	}

	prompt := os.Args[1]

	api_key := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(api_key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,  //"Hello!"
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}


