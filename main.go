package main

import (
	"MatchZy-Webhook/util"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gtuk/discordwebhook"
	"github.com/wizzard0/trycloudflared"
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
		Port uint16 `json:"port"`
		Auth struct {
			Header string `json:"header"`
			Value  string `json:"value"`
		} `json:"auth"`
	} `json:"server"`
	Discord struct {
		AvatarUrl string `json:"avatar_url"`
		Webhook   string `json:"webhook"`
	} `json:"discord"`
	Cloudflare struct {
		UseCloudflare bool     `json:"use_cloudflare"`
		FreeTunnel    bool     `json:"free_tunnel"`
		Account       struct{} `json:"account"`
	} `json:"cloudflare"`
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

	if cfg.Cloudflare.UseCloudflare {
		if cfg.Cloudflare.FreeTunnel {
			ctx, cancel := context.WithCancel(context.Background())

			url, err := trycloudflared.CreateCloudflareTunnel(ctx, int(cfg.Server.Port))
			if err != nil {
				log.Println("Error creating tunnel. ", err)
			}

			log.Println("Tunnel url: ", url)

			log.Printf("Enter these commands in console: \n matchzy_remote_log_url %s/matchzy \n matchzy_remote_log_header_key %s \n matchzy_remote_log_header_value %s", url, cfg.Server.Auth.Header, cfg.Server.Auth.Value)

			defer cancel()
		}
	}

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
	log.Println("event: ", res.Event)

	msg = replaceMsg(msg, res)

	log.Println("msg: ", msg)

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
	return processForeach(msg, reflect.ValueOf(data))
}
func processForeach(msg string, v reflect.Value) string {
	for {
		start := strings.Index(msg, "$FOREACH(")
		if start == -1 {
			break
		}
		end := strings.Index(msg[start:], ")")
		if end == -1 {
			break
		}
		end += start
		placeholder := msg[start : end+1]
		re := regexp.MustCompile(`\$FOREACH\(\$(\w+),\s*'([^']*)'\)`)
		matches := re.FindStringSubmatch(placeholder)
		if len(matches) != 3 {
			break
		}
		tag := matches[1]
		innerTemplate := matches[2]
		sliceVal := findSliceByTag(v, tag)
		if !sliceVal.IsValid() || sliceVal.Len() == 0 {
			msg = strings.ReplaceAll(msg, placeholder, "")
			continue
		}
		var repeated string
		for i := 0; i < sliceVal.Len(); i++ {
			repeated += walk(innerTemplate, sliceVal.Index(i))
		}
		msg = strings.ReplaceAll(msg, placeholder, repeated)
	}
	return walk(msg, v)
}
func walk(msg string, v reflect.Value) string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return msg
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("str")
		if tag != "" {
			msg = strings.ReplaceAll(msg, "$"+tag, valueToString(field))
		}
		if field.Kind() == reflect.Struct {
			msg = walk(msg, field)
		}
	}
	return msg
}
func findSliceByTag(v reflect.Value, tag string) reflect.Value {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.Value{}
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldTag := t.Field(i).Tag.Get("str")
		if field.Kind() == reflect.Slice && fieldTag == tag {
			return field
		}
		if field.Kind() == reflect.Struct {
			if res := findSliceByTag(field, tag); res.IsValid() {
				return res
			}
		}
	}
	return reflect.Value{}
}
func valueToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprint(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprint(v.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprint(v.Float())
	case reflect.Bool:
		return fmt.Sprint(v.Bool())
	}
	return ""
}
