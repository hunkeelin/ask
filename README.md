# Ask - AI Assistant CLI

A simple command-line AI assistant that supports both OpenAI and Google Gemini APIs. Ask questions and get responses from your preferred AI service directly from your terminal.

## Features

- **Multi-provider support**: Choose between OpenAI (GPT-3.5-turbo) and Google Gemini
- **Simple CLI interface**: Just type your question as command arguments
- **Environment-based configuration**: Easy setup with environment variables
- **Built-in help**: Comprehensive help message with usage examples
- **No quotes required**: Ask questions naturally without wrapping in quotes

## Installation

1. Clone or download this repository
2. Make sure you have Go installed (version 1.18 or later)
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Setup

### Environment Variables

You need to set up the following environment variables:

#### Required
- **`AI_PROVIDER`** - Choose which AI service to use:
  - `openai` - Use OpenAI's GPT-3.5-turbo
  - `gemini` - Use Google's Gemini (default if not set)

#### API Keys (one required based on provider)
- **`OPENAI_API_KEY`** - Required when using OpenAI
- **`GEMINI_API_KEY`** - Required when using Gemini

### Getting API Keys

#### OpenAI API Key
1. Go to [OpenAI Platform](https://platform.openai.com/)
2. Sign in or create an account
3. Navigate to API Keys section
4. Create a new API key
5. Copy and save it securely

#### Gemini API Key
1. Go to [Google AI Studio](https://aistudio.google.com/)
2. Sign in with your Google account
3. Click "Get API key"
4. Create a new API key
5. Copy and save it securely

### Environment Setup Examples

#### Using OpenAI
```bash
export AI_PROVIDER=openai
export OPENAI_API_KEY=your_openai_api_key_here
```

#### Using Gemini (default)
```bash
export AI_PROVIDER=gemini  # Optional since it's the default
export GEMINI_API_KEY=your_gemini_api_key_here
```

#### Persistent Setup
Add to your shell profile (`.bashrc`, `.zshrc`, etc.):
```bash
# Choose your preferred provider
export AI_PROVIDER=gemini
export GEMINI_API_KEY=your_gemini_api_key_here

# Or for OpenAI
# export AI_PROVIDER=openai
# export OPENAI_API_KEY=your_openai_api_key_here
```

## Usage

### Basic Usage
```bash
go run main.go your question here
```

### Examples
```bash
# Ask about the weather
go run main.go what is the weather like today

# Get coding help
go run main.go explain how to use goroutines in Go

# Ask for explanations
go run main.go explain quantum computing in simple terms

# Creative requests
go run main.go write a short poem about programming
```

### Switching Providers
```bash
# Use OpenAI for this session
export AI_PROVIDER=openai
go run main.go explain machine learning

# Switch to Gemini
export AI_PROVIDER=gemini
go run main.go what are the benefits of renewable energy
```

### Help
```bash
go run main.go --help
# or
go run main.go -h
```

## Building

To build a standalone executable:

```bash
go build -o ask main.go
```

Then run directly:
```bash
./ask your question here
```

## Dependencies

This project uses the following Go modules:
- `github.com/google/generative-ai-go` - Google Gemini API client
- `github.com/sashabaranov/go-openai` - OpenAI API client
- `google.golang.org/api` - Google API client library

## Project Structure

```
ask/
├── main.go          # Main application code
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
└── README.md        # This file
```

## Error Handling

The application provides clear error messages for common issues:

- **Missing prompt**: "Error: No prompt provided"
- **Missing API key**: Specific message based on selected provider
- **Invalid provider**: Lists valid options (openai, gemini)
- **API errors**: Displays the actual API error message

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## Security Notes

- **Never commit API keys** to version control
- Store API keys as environment variables
- Consider using tools like `direnv` for project-specific environment variables
- Rotate API keys regularly

## Troubleshooting

### Common Issues

**"No API key found" error**
- Make sure you've set the correct environment variable for your chosen provider
- Verify the API key is valid and hasn't expired

**"Invalid AI_PROVIDER" error**
- Check that `AI_PROVIDER` is set to either `openai` or `gemini`
- The value is case-insensitive

**API rate limit errors**
- Wait a moment and try again
- Check your API usage limits on the respective platforms

### Debug Mode
To see more detailed error information, you can modify the log level or add debug prints to the source code.

## License

This project is provided as-is for educational and personal use.

## Changelog

### v1.0.0
- Initial release
- Support for OpenAI GPT-3.5-turbo
- Support for Google Gemini
- Environment-based provider selection
- Built-in help system
