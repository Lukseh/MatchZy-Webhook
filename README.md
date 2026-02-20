# MatchZy-Webhook

Server listening for MatchZy events and formatting them to display using Discord webhooks.

## Setup

- copy config.example.json to config.json and fill it out

- use ``` go run main.go ``` or build it into .exe with ``` go build main.go ```

- Edit templates.json to your needs

- Set matchzy_remote_log_url to $Your_domain/matchzy and matchzy_remote_log_header_value to same value as(AUTH) from env

## Contribution

Everyone is open to contribute with either making types for Events or adding str tags for $VAR templating.
