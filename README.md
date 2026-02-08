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
After sending, the file ID will be displayed and automatically replied to the message in Telegram.

### Send a file with caption:
```bash
MdsFileShare --chat "-143242342" --file "document.pdf" --caption "Important document"
```

### Download a file using file ID:
```bash
MdsFileShare --chat "-143242342" --getfile "BQACAgQAAxkBAAIC..."
```

### Download with custom output path:
```bash
MdsFileShare --chat "-143242342" --getfile "BQACAgQAAxkBAAIC..." --output "downloads/myfile.pdf"
```

### Find chat ID:
```bash
MdsFileShare --find "31311"
```
Then send "31311" as a message from any chat/group/channel where the bot is added. The tool will display the chat name and ID.

## Options

- `--chat` (required for sending/downloading): Telegram chat ID
- `--file` (required for sending): Path to the file to send
- `--caption` (optional): Caption text for the file
- `--getfile` (alternative mode): Download file using Telegram file ID
- `--output` (optional): Custom output path for downloaded file
- `--find` (alternative mode): Listen for messages containing this text and return chat info

## Features

- ✅ Cross-platform (Windows, macOS, Linux)
- 📊 Real-time upload/download progress bar
- 📁 File information display (name, size, type)
- 📤 Send files to Telegram chats
- 📥 Download files using file ID
- 📎 Automatic file ID reply after upload
- 🔍 Find chat IDs by sending a unique text
- 🎯 Simple command-line interface
- 🚀 Single binary executable
- 📱 Works with private chats, groups, supergroups, and channels
