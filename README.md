# MatchZy-Webhook

Server listening for MatchZy events and formatting them to display using Discord webhooks.

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
