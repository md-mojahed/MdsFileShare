# MdsFileShare

A cross-platform command-line tool to send files to Telegram chats with progress tracking.

## Installation

```bash
go mod download
go build -o MdsFileShare
```

## Usage

### Send a file:
```bash
MdsFileShare --chat "-143242342" --file "path/to/file.pdf"
```

### Send a file with caption:
```bash
MdsFileShare --chat "-143242342" --file "document.pdf" --caption "Important document"
```

### Find chat ID:
```bash
MdsFileShare --find "31311"
```
Then send "31311" as a message from any chat/group/channel where the bot is added. The tool will display the chat name and ID.

## Options

- `--chat` (required for sending): Telegram chat ID
- `--file` (required for sending): Path to the file to send
- `--caption` (optional): Caption text for the file
- `--find` (alternative mode): Listen for messages containing this text and return chat info

## Features

- ✅ Cross-platform (Windows, macOS, Linux)
- 📊 Real-time upload progress bar
- 📁 File information display (name, size, type)
- 🔍 Find chat IDs by sending a unique text
- 🎯 Simple command-line interface
- 🚀 Single binary executable
- 📱 Works with private chats, groups, supergroups, and channels
