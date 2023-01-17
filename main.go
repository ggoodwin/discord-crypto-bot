package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/octolibs/stocks"
)

/** Globals
 * * Global Variables
 */
var (
	DiscordToken string
	CryptoTicker string
	config       *Config
)

/** Config Struct
 * * This is returned by our config file
 */
type Config struct {
	DiscordToken string `json:"DiscordToken"`
	CryptoTicker string `json:"CryptoTicker"`
}

/** ReadConfig
 * * Reads our config file and assigns values to globals
 */
func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := os.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	DiscordToken = config.DiscordToken
	CryptoTicker = config.CryptoTicker

	return nil

}

/** Start
* * Main Worker Function
 */
func Start() {
	var goBot *discordgo.Session
	var BotId string

	fmt.Println("Starting...")
	goBot, err := discordgo.New(fmt.Sprintf("Bot %s", DiscordToken))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Bot %s is running !", BotId)

	goBot.UpdateGameStatus(0, createMessage())

	ticker := time.Tick(time.Minute)

	for range ticker {
		goBot.UpdateGameStatus(0, createMessage())
	}
}

/** createMessage
 * * Creates the actual playing message
 */
func createMessage() string {
	price, percent, direction := stocks.GetPriceAndPercentage(CryptoTicker)
	return fmt.Sprintf("%s %s %s", price, direction, percent)
}

/**
* Main
* * Main Injection Point
 */
func main() {
	err := ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	<-make(chan struct{})
}
