# go-gpt

[OpenAI GPT-3](https://beta.openai.com/) API wrapper for Go
[#官方文档](https://platform.openai.com/docs/introduction)
[#例子](https://platform.openai.com/examples)
[#官方playground](https://platform.openai.com/playground)

Installation:
```
go get github.com/vurtneyang/go-gpt
```


Example usage:

```go
package main

import (
	"context"
	"fmt"
	gogpt "github.com/vurtneyang/go-gpt"
)

func main() {
	c := gogpt.NewClient("your token")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
```
