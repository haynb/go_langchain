package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"os"
)

func main() {
	if err := os.Setenv("OPENAI_API_KEY", "sk-U*************************iXGPa9r"); err != nil {
		panic(err)
	}

	if err := os.Setenv("OPENAI_BASE_URL", "https://cfwus02.opapi.win/v1/"); err != nil {
		panic(err)
	}
	if err := os.Setenv("OPENAI_MODEL", "gpt-3.5-turbo"); err != nil {
		panic(err)
	}
	llm, err := openai.NewChat()
	if err != nil {
		panic(err)
	}
	completion, err := llm.Call(context.Background(), []schema.ChatMessage{
		//schema.SystemChatMessage{Content: "Hello, I am a friendly chatbot. I love to talk about movies, books and music. Answer in long form yaml."},
		schema.HumanChatMessage{Content: "鲁迅为什么要打周树人?"},
	}, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Println("chunk: ", string(chunk)) // 流式输出
		return nil
	}),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("completion：", completion)
}
