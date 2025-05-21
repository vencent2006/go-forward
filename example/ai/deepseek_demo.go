package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// 定义请求和响应的结构体
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int     `json:"index"`
		Message      Message `json:"message"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func main() {
	apiKey := os.Getenv("DEEPSEEK_API_KEY")                  // 替换为你的DeepSeek API key
	apiUrl := "https://api.deepseek.com/v1/chat/completions" // DeepSeek API端点

	// 创建请求体
	requestBody := ChatRequest{
		Model: "deepseek-chat", // 指定模型
		Messages: []Message{
			{
				Role:    "user",
				Content: "今天星期几",
			},
		},
		Stream: false,
	}

	// 将请求体编码为JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error marshaling request body: %v\n", err)
		return
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status code %d: %s\n", resp.StatusCode, string(body))
		return
	}

	// 解析响应
	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return
	}
	fmt.Printf("Response: %+v\n", chatResponse)

	// 输出结果
	if len(chatResponse.Choices) > 0 {
		fmt.Println("AI回复:", chatResponse.Choices[0].Message.Content)
	} else {
		fmt.Println("没有收到有效的回复")
	}
}
