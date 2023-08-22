package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/textsplitter"
	"os"
)

func main() {
	// 定义要设置的环境变量和对应的值
	envVars := map[string]string{
		"OPENAI_API_KEY":  "sk-UhzxW8avagN8D00Q2AmHeJL1LBz6NFv9rI3USa94GiXGPa9r",
		"OPENAI_BASE_URL": "https://cfwus02.opapi.win/v1/",
		"OPENAI_MODEL":    "gpt-3.5-turbo",
	}

	// 逐个设置环境变量
	for key, value := range envVars {
		err := os.Setenv(key, value)
		if err != nil {
			panic(err)
		}
	}
	llm, err := openai.NewChat()
	if err != nil {
		panic(err)
	}
	stuffQAChain := chains.LoadStuffQA(llm)
	file, err := os.Open("xiaozhao.pdf")
	if err != nil {
		panic(err)
	}
	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fileSize := fileInfo.Size()
	p := documentloaders.NewPDF(file, fileSize)
	spliter := textsplitter.NewTokenSplitter()
	doc, err := p.LoadAndSplit(context.Background(), spliter)
	if err != nil {
		panic(err)
	}
	answer, err := chains.Call(context.Background(), stuffQAChain, map[string]any{
		"input_documents": doc,
		"question":        "住宿费多少钱?",
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(answer)
}
