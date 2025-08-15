package main

import (
  "context"
  "flag"
  "fmt"
  "os"
  "strings"

  "github.com/google/generative-ai-go/genai"
  "google.golang.org/api/option"
  "github.com/sashabaranov/go-openai"
)

func main() {
  // Set up command line flags
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "AI Assistant - Chat with OpenAI or Gemini\n\n")
    fmt.Fprintf(os.Stderr, "Usage: %s [options] your prompt here\n\n", os.Args[0])
    fmt.Fprintf(os.Stderr, "Environment Variables:\n")
    fmt.Fprintf(os.Stderr, "  AI_PROVIDER      - Choose which AI service to use: 'openai' or 'gemini' (default: gemini)\n")
    fmt.Fprintf(os.Stderr, "  OPENAI_API_KEY   - Required when using OpenAI\n")
    fmt.Fprintf(os.Stderr, "  GEMINI_API_KEY   - Required when using Gemini\n\n")
    fmt.Fprintf(os.Stderr, "Examples:\n")
    fmt.Fprintf(os.Stderr, "  export AI_PROVIDER=openai && %s what is the weather today\n", os.Args[0])
    fmt.Fprintf(os.Stderr, "  export AI_PROVIDER=gemini && %s explain quantum computing\n\n", os.Args[0])
    fmt.Fprintf(os.Stderr, "Options:\n")
    flag.PrintDefaults()
  }

  flag.Parse()

  // Check if a prompt was provided as a command-line argument.
  if flag.NArg() == 0 {
    fmt.Fprintf(os.Stderr, "Error: No prompt provided\n\n")
    flag.Usage()
    os.Exit(1)
  }

  // Join all remaining arguments to form a single prompt.
  prompt := strings.Join(flag.Args(), " ")

  // Set up the context with a timeout. This is good practice for network requests.
  ctx := context.Background()

  // Get the AI provider choice and API keys
  provider := strings.ToLower(os.Getenv("AI_PROVIDER"))
  if provider == "" {
    provider = "gemini" // Default to Gemini
  }

  openaiKey := os.Getenv("OPENAI_API_KEY")
  geminiKey := os.Getenv("GEMINI_API_KEY")

  switch provider {
  case "openai":
    if openaiKey == "" {
      panic("OPENAI_API_KEY environment variable not set. Required when AI_PROVIDER=openai")
    }
    generateWithOpenAI(ctx, openaiKey, prompt)
  case "gemini":
    if geminiKey == "" {
      panic("GEMINI_API_KEY environment variable not set. Required when AI_PROVIDER=gemini")
    }
    generateWithGemini(ctx, geminiKey, prompt)
  default:
    panic(fmt.Sprintf("Invalid AI_PROVIDER '%s'. Must be 'openai' or 'gemini'", provider))
  }
}

func generateWithOpenAI(ctx context.Context, apiKey, prompt string) {
  client := openai.NewClient(apiKey)

  resp, err := client.CreateChatCompletion(
    ctx,
    openai.ChatCompletionRequest{
      Model: openai.GPT3Dot5Turbo,
      Messages: []openai.ChatCompletionMessage{
        {
          Role:    openai.ChatMessageRoleUser,
          Content: prompt,
        },
      },
    },
  )

  if err != nil {
    panic(err)
  }

  fmt.Println("\nOpenAI:")
  fmt.Println(resp.Choices[0].Message.Content)
}

func generateWithGemini(ctx context.Context, apiKey, prompt string) {
  // Create a new client.
  client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
  if err != nil {
    panic(err)
  }
  defer client.Close()

  // Create the model instance.
  model := client.GenerativeModel("gemini-1.5-flash")

  // Generate content from the user's prompt.
  resp, err := model.GenerateContent(ctx, genai.Text(prompt))
  if err != nil {
    panic(err)
  }

  // Print the generated content.
  fmt.Println("\nGemini:")
  for _, part := range resp.Candidates[0].Content.Parts {
    fmt.Println(part)
  }
}
