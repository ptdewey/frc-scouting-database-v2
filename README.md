# FRC Scouting Database V2

This project exists to ease the burden on human scouts during busy and exciting robotics competitions.
The event analyzer utilizes data from the Blue Alliance API to fetch team and match data from targeted events.

Version 2 iterates on the previous version by introducing automation, more features, and cleaner code.
V2 is written entirely in Go, enabling more consistent and faster execution, and to maintain better tests and reproducibility from the original version.

## How to Use:

1. Ensure [dependencies](#Dependencies) are installed.
2. Clone the repository `git clone https://github.com/ptdewey/frc-scouting-database-v2.git`
```bash
git clone https://github.com/ptdewey/frc-scouting-database-v2.git
```
3. Add 2 .env files in `config/` (create this directory if necessary):
    - app.env (containing a The Blue Alliance api key that can be obtained from their website)
        - `API_KEY="{api-key}"`
    - bot.env (containg a discord bot token and channel id to send automated updates to)
        - `DISCORD_BOT_TOKEN="{bot-token}"`
        - `DISCORD_CHANNEL_ID="{channl-id}"`
4. Run using Docker Compose.
```bash
docker-compose up --build -d
```

After this, the application will be running in two separate containers, one for the data exporter bot and one for the event analyzer.
Output data can be found in the project-root/output directory with subfolders for years and events, with event folders being named by event key.

<!-- TODO: data dictionary -->

## Dependencies

| Name              | Version    |
| ------------------|------------|
| Docker            | >= 24.0.5  |
| docker-compose    | >= 1.29.2  |


---

Powered by [The Blue Alliance](https://thebluealliance.com/)
