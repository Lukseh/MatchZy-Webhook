# MatchZy-Webhook

MatchZy Webhook listens for Events on your MatchZy server and pushes Discord message formatted via templates allowing for Customisation with option to display data from server with placeholders like $MAP etc...

List of these is available [here](https://github.com/Lukseh/MatchZy-Webhook/wiki/Templating)

Support original plugin and its creator [here](https://github.com/shobhit-pathak/MatchZy)

Supports all Events with respective templates.

To use `$DUROUND` template please set the value of `"round_time"` in `config.json` to same value as it is in `game/csgo/cfg/MatchZy/live.cfg` or it will not display clock properly.

Config allows for using free Cloudflared tunnel to securely expose the server fast and easy and gives copy & paste commands in terminal allowing for quick setup before matches.

## Setup

### Manual build

- copy config.example.json to config.json and fill it out

- use ``` go run main.go ``` or build it into .exe with ``` go build main.go ```

- Edit templates.json to your needs

- Set matchzy_remote_log_url to $Your_domain/matchzy and matchzy_remote_log_header_value to same value as(AUTH) from env

### Docker

I included Dockerfile and docker-compose.yaml and made automatically updated image ghcr.io/lukseh/matchzy-webhook:main so you can see how to setup fast and easy.

Please use Nginx or any other proxy in front of it, or at least use VPS with good firewall and fail2ban if using opened ports

I included the example with external network so you can uncomment part that you need and change name to connect it to your proxy.

## Contribution

Everyone is open to contribute with either making types for Events or adding str tags for $VAR templating.
