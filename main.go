package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "!gob"

func main() {
	// export DISCORD_BOT_TOKEN=blah
	sess, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		args := strings.Split(m.Content, " ")
		if args[0] != prefix {
			return
		} else if args[1] == "proverbs" {
			proverbs := []string{
				"Don't communicate by sharing memory, share memory by communicating.",
				"Concurrency is not parallelism.",
				"Channels orchestrate; mutexes serialize.",
				"The bigger the interface, the weaker the abstraction.",
				"Make the zero value useful.",
				"interface{} says nothing.",
				"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
			}
			selection := rand.Intn(len(proverbs))
			author := discordgo.MessageEmbedAuthor{
				Name: "Rob Pike",
				URL:  "https://go-proverbs.github.io/",
			}
			embed := discordgo.MessageEmbed{
				Title:  proverbs[selection],
				Author: &author,
			}
			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			//s.ChannelMessageSend(m.ChannelID, proverbs[selection])
		}
		if args[1] == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world")
		}
	})
	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("ITS ON")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	for {
		<-sc
		fmt.Println("Terminating bot...")
		return
	}

}
