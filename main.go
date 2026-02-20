package main

import (
	"MatchZy-Webhook/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gtuk/discordwebhook"
	_ "github.com/joho/godotenv/autoload" // lazy load env
)

type info struct {
	username string
	avatar   string
	url      string
}
type templates map[util.Event]string

type config struct {
	Server struct {
		Host string `json:"host"`
		Port uint8  `json:"port"`
		Auth string `json:"auth"`
	} `json:"server"`
	Discord struct {
		AvatarUrl string `json:"avatar_url"`
		Webhook   string `json:"webhook"`
	} `json:"discord"`
}

func main() {

	var cfg config
	configFile, err := os.ReadFile("./config/config.json")
	if err != nil {
		log.Println("Error parsing config: ", err)
	}
	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		log.Println("Error parsing config: ", err)
	}

	info := info{
		username: "MatchZy Webhooks",
		avatar:   cfg.Discord.AvatarUrl,
		url:      cfg.Discord.Webhook,
	}

	// Logging to file
	var logFile *os.File
	t := time.Now()
	formatted := t.Format("2006-01-02-15:04:05")
	if err := os.MkdirAll("./logs", 0755); err != nil {
		log.Fatalln("Failed to create logs directory:", err)
	}
	if logFile, err = os.Create(fmt.Sprint("./logs/", formatted, ".log")); err != nil {
		log.Fatalln("Error opening log file.\n ERR: ", err.Error())
	}
	mw := io.MultiWriter(log.Writer(), logFile)
	log.SetOutput(mw)

	app := fiber.New()

	// Templates
	var templates templates
	UpdateTemplates := func() {
		templatefile, err := os.ReadFile("templates.json")
		if err != nil {

		}
		err = json.Unmarshal(templatefile, &templates)
		if err != nil {

		}
	}

	UpdateTemplates() // It needs to be outside function because i want to add watching for changes on template file.

	app.Use("/", static.New("./public/"))
	app.Post("/matchzy", util.AuthCheck(cfg.Server.Auth), func(c fiber.Ctx) error {
		SendMsg(templates, &info, c.BodyRaw())
		return c.Status(fiber.StatusOK).SendString("MatchZy endpoint.")
	})
	log.Printf("Listening on http://%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Fatalln(app.Listen(fmt.Sprint(cfg.Server.Host, ":", cfg.Server.Port), fiber.ListenConfig{
		DisableStartupMessage: true,
	}))
}

func SendMsg(templates templates, info *info, msgraw []byte) {

	var res util.MatchZyRes
	if err := json.Unmarshal(msgraw, &res); err != nil {
		log.Println("JSON error:", err)
		return
	}
	var msg string

	msg = templates[res.Event]

	msg = replaceMsg(msg, res)

	err := discordwebhook.SendMessage(info.url, discordwebhook.Message{
		Username:  &info.username,
		AvatarUrl: &info.avatar,
		Content:   &msg,
	})
	if err != nil {
		log.Println(err)
	}
}

func replaceMsg(msg string, data any) string {
	return walk(msg, reflect.ValueOf(data))
}

// I have no idea about anything with tags so this is vibe coded (it works for now)
func walk(msg string, v reflect.Value) string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return msg
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("str")
		if tag != "" {
			var replacement string
			switch fieldValue.Kind() {
			case reflect.String:
				replacement = fieldValue.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				replacement = fmt.Sprint(fieldValue.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				replacement = fmt.Sprint(fieldValue.Uint())
			case reflect.Float32, reflect.Float64:
				replacement = fmt.Sprint(fieldValue.Float())
			case reflect.Bool:
				replacement = fmt.Sprint(fieldValue.Bool())
			}
			msg = strings.ReplaceAll(msg, "$"+tag, replacement)
		}
		if fieldValue.Kind() == reflect.Struct {
			msg = walk(msg, fieldValue)
		}
	}

	return msg
}
