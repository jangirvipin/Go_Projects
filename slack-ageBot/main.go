package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"sync"
)

func printEvents(events <-chan *slacker.CommandEvent, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Command Events:")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping command event listener")
			return
		case event, ok := <-events:
			if !ok {
				fmt.Println("Command event channel closed")
				return
			}
			fmt.Println(event.Timestamp)
			fmt.Println(event.Command)
			fmt.Println(event.Parameters)
			fmt.Println(event.Event)
			fmt.Println()
		}
	}
}

func main() {
	var wg sync.WaitGroup

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go printEvents(bot.CommandEvents(), ctx, &wg)

	var example = []string{"My name is John", "My name is Vipin"}

	bot.Command("My name is <name>", &slacker.CommandDefinition{
		Description: "Greet the user by name",
		Examples:    example,
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			name := request.Param("name")
			if name == "" {
				_ = response.Reply("Name is required")
				return
			}
			err := response.Reply("Hello" + name + ", nice to meet you!")
			if err != nil {
				return
			}
		},
	})

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
	wg.Wait()
}
