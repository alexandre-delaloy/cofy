package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"github.com/blyndusk/cofy/app/commands"
)

func init() {
	// load environment variables, return error if doesn't exist
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// get Discord token
	Token := os.Getenv("BOT_TOKEN")

	// create a variable named "t", as the bot token
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	setupSession()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := "c!"

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	commands.Info(s, m, prefix)
}

func reactToCoffee(s *discordgo.Session, m *discordgo.MessageCreate) {

	matched, err := regexp.MatchString(`coffee|coffe|cofee|cofe|cafe|café`, m.Content)
	fmt.Println(matched, err)

	if matched == true {
		if m.Author.ID == s.State.User.ID {
			return
		}
		s.ChannelMessageSend(m.ChannelID, "...mMmMMmmMmmmHhHHhhHhhh... ...cOoOOooOoooffEeEEeeEeee...")
		s.MessageReactionAdd(m.ChannelID, m.ID, "☕")

	}

}

func setupSession() {
	Token := os.Getenv("BOT_TOKEN")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(reactToCoffee)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
