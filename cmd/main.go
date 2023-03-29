package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {

	err := godotenv.Load("configs/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := openai.NewClient(os.Getenv("API_KEY"))
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Hola, ¿cómo estás?",
			},
		},
	}

	fmt.Println(req.Messages[0].Content + "\n")

	for {
		resp, err := client.CreateChatCompletion(
			ctx,
			req,
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		req.Messages[0].Role = openai.ChatMessageRoleUser
		req.Messages[0].Content = resp.Choices[0].Message.Content

		fmt.Println(resp.Choices[0].Message.Content + "\n")
	}
}
