package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/blyndusk/cofy/app/commands"
	"github.com/blyndusk/cofy/app/services"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func init() {
	// get Discord token
	Token := services.EnvVar("BOT_TOKEN")

	// create a variable named "t", as the bot token
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	pingApi()
	setupSession()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	commands.SetCommands(s, m)
	// commands.Profile(s, m, prefix)
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
	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func pingApi() {
	_, err := http.Get(services.EnvVar("API_URL"))
	if err != nil {
		log.Fatal(err)
	}
	log.Info("API available !")
}
