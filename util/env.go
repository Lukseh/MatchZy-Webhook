package util

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func ParseConfig(domain *string, host *string, port *string, avatar *string, auth *string, discordurl *string) {
	// Config
	*domain = os.Getenv("DOMAIN")
	*host = os.Getenv("HOST")
	*port = os.Getenv("PORT")
	if *port == "" {
		*port = "74"
	}

	*auth = os.Getenv("AUTH")
	if *auth == "" {
		log.Fatalln("Empty AUTH header.")
	}
	*domain = os.Getenv("DOMAIN")

	*avatar = os.Getenv("AVATAR_URL")
	if *avatar != "" {
		request := fiber.Get(*avatar)
		request.Debug()
		if status, _, _ := request.Bytes(); status != 200 {
			log.Println("Avatar url does not exist, using default.")
			*avatar = fmt.Sprint(*domain, "/def.png")
		}
	} else {
		*avatar = fmt.Sprint(*domain, "/def.png")
	}

	*discordurl = os.Getenv("DISCORD_WB")
	if *discordurl == "" {
		log.Fatalln("Did not provide webhook url.")
	}
}
