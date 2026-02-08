package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/schollz/progressbar/v3"
)

const botToken = "8542820282:AAGuuBZOGjulWPPA5xYjczlql7hiVKrC8rU"

func main() {
	chatID := flag.String("chat", "", "Telegram chat ID (required for file sending)")
	filePath := flag.String("file", "", "Path to file to send (required for file sending)")
	caption := flag.String("caption", "", "Optional caption for the file")
	findText := flag.String("find", "", "Listen for messages containing this text and return chat info")
	flag.Parse()

	// Find mode
	if *findText != "" {
		if err := findChat(*findText); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Send file mode
	if *chatID == "" || *filePath == "" {
		fmt.Println("Error: --chat and --file are required for sending files")
		fmt.Println("Or use --find \"text\" to find chat IDs")
		flag.Usage()
		os.Exit(1)
	}

	if err := sendFile(*chatID, *filePath, *caption); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func sendFile(chatIDStr, filePath, caption string) error {
	// Parse chat ID
	var chatID int64
	if _, err := fmt.Sscanf(chatIDStr, "%d", &chatID); err != nil {
		return fmt.Errorf("invalid chat ID: %v", err)
	}

	// Check if file exists
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("file not found: %v", err)
	}

	// Display file information
	fmt.Println("\n📁 File Information:")
	fmt.Printf("   Name: %s\n", fileInfo.Name())
	fmt.Printf("   Size: %s\n", formatSize(fileInfo.Size()))
	fmt.Printf("   Type: %s\n", getFileType(filePath))
	fmt.Printf("   Chat: %s\n", chatIDStr)
	if caption != "" {
		fmt.Printf("   Caption: %s\n", caption)
	}
	fmt.Println()

	// Initialize bot
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return fmt.Errorf("failed to initialize bot: %v", err)
	}

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create progress bar
	bar := progressbar.NewOptions64(
		fileInfo.Size(),
		progressbar.OptionSetDescription("📤 Uploading"),
		progressbar.OptionSetWidth(40),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Println()
		}),
	)

	// Create progress reader
	reader := &progressReader{
		reader: file,
		bar:    bar,
	}

	// Prepare document message
	doc := tgbotapi.NewDocument(chatID, tgbotapi.FileReader{
		Name:   filepath.Base(filePath),
		Reader: reader,
	})

	if caption != "" {
		doc.Caption = caption
	}

	// Send file
	_, err = bot.Send(doc)
	if err != nil {
		return fmt.Errorf("failed to send file: %v", err)
	}

	fmt.Println("✅ File sent successfully!")
	return nil
}

// progressReader wraps an io.Reader to update progress bar
type progressReader struct {
	reader *os.File
	bar    *progressbar.ProgressBar
}

func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.bar.Add(n)
	return n, err
}

// formatSize converts bytes to human-readable format
func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// getFileType returns the file extension or type
func getFileType(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext == "" {
		return "Unknown"
	}
	return ext[1:] // Remove the dot
}

// findChat listens for messages and finds chats that send the specified text
func findChat(searchText string) error {
	fmt.Printf("🔍 Listening for messages containing: \"%s\"\n", searchText)
	fmt.Println("📱 Send this text from any chat/group/channel where the bot is added...")
	fmt.Println("⏳ Waiting for messages... (Press Ctrl+C to stop)\n")

	// Initialize bot
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return fmt.Errorf("failed to initialize bot: %v", err)
	}

	bot.Debug = false

	// Get bot info
	botInfo, err := bot.GetMe()
	if err != nil {
		return fmt.Errorf("failed to get bot info: %v", err)
	}
	fmt.Printf("✅ Bot connected: @%s\n\n", botInfo.UserName)

	// Configure updates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Listen for updates
	for update := range updates {
		if update.Message == nil {
			continue
		}

		message := update.Message
		
		// Check if message text contains the search text
		if strings.Contains(message.Text, searchText) {
			chatType := getChatType(message.Chat.Type)
			chatName := getChatName(message.Chat)
			
			fmt.Println("✨ Match found!")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Printf("📋 Chat Name: %s\n", chatName)
			fmt.Printf("🆔 Chat ID: %d\n", message.Chat.ID)
			fmt.Printf("📱 Chat Type: %s\n", chatType)
			fmt.Printf("👤 Sender: %s\n", getSenderName(message.From))
			fmt.Printf("💬 Message: %s\n", message.Text)
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Printf("\n✅ Use this chat ID: %d\n\n", message.Chat.ID)
			fmt.Println("🔍 Still listening for more matches... (Press Ctrl+C to stop)")
		}
	}

	return nil
}

// getChatType returns a human-readable chat type
func getChatType(chatType string) string {
	switch chatType {
	case "private":
		return "Private Chat"
	case "group":
		return "Group"
	case "supergroup":
		return "Supergroup"
	case "channel":
		return "Channel"
	default:
		return chatType
	}
}

// getChatName returns the chat name or title
func getChatName(chat *tgbotapi.Chat) string {
	if chat.Title != "" {
		return chat.Title
	}
	if chat.UserName != "" {
		return "@" + chat.UserName
	}
	if chat.FirstName != "" {
		name := chat.FirstName
		if chat.LastName != "" {
			name += " " + chat.LastName
		}
		return name
	}
	return "Unknown"
}

// getSenderName returns the sender's name
func getSenderName(user *tgbotapi.User) string {
	if user == nil {
		return "Unknown"
	}
	name := user.FirstName
	if user.LastName != "" {
		name += " " + user.LastName
	}
	if user.UserName != "" {
		name += " (@" + user.UserName + ")"
	}
	return name
}
