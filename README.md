# FastWand 🪄

A magical CLI tool to initialize FastHTML + Tailwind + DaisyUI projects. FastWand automatically downloads and installs the Tailwind CSS CLI (optionally bundled with DaisyUI) for your operating system and architecture.

⚠️ **ALPHA STATUS**: This project is in early development. Expect bugs and breaking changes.

## Features

- 🎨 Choice between vanilla Tailwind CSS or Tailwind + DaisyUI
- 🚀 Automatic binary downloads for your system
- 💅 Beautiful TUI with Bubble Tea
- 🔄 Live reload during development
- 📦 Zero-config setup

## Installation

### Via Python Package (Recommended)
```bash
pip install fastwand
```

### Direct Binary Download
Download the latest binary for your system from our [releases page].

## Usage

### Initialize a New Project
```bash
fastwand init [directory]
```

This will:
- Present a beautiful UI to choose your framework
- Download the appropriate Tailwind CLI for your system
- Create a basic FastHTML project structure
- Set up Tailwind CSS configuration

### Development Mode
Run these commands in separate terminal windows:
```bash
# Terminal 1: Watch for CSS changes
fastwand watch

# Terminal 2: Run the Python server
python main.py
```

### Production Mode
Build the minified CSS and run the server in one command:
```bash
fastwand run
```

## System Compatibility

Automatically detects and supports:
- Linux (x64, arm64)
- macOS (x64, arm64)
- Windows (x64)

## Development

### Prerequisites
- Go 1.22.5 or higher

### Building from Source
```bash
go build
```

### Running Tests
```bash
go test ./...
```

## Contributing

This project is in active development. Issues and pull requests are welcome!

## Acknowledgements

FastWand builds upon and is inspired by several excellent projects:

- [FastHTML](https://github.com/answerDotAI/fasthtml/) - The core HTML framework that FastWand is built to support
- [Tailwind CSS](https://github.com/tailwindlabs/tailwindcss) - The utility-first CSS framework that powers our styling
- [Tailwind CLI Extra](https://github.com/dobicinaitis/tailwind-cli-extra) - Pre-bundled Tailwind CLI with DaisyUI plugin, which we use for binary distribution


## License

MIT