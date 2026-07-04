# IP9
A dead simple Golang program that returns your IP address and country.

Country detection uses [ip2region](https://github.com/lionsoul2014/ip2region), a free and open-source offline IP geolocation library (Apache 2.0). No API keys or external services are required at runtime.

Example response:

```
203.0.113.42
Australia (AU)
```

Build and run with docker-compose:

```
docker compose up -d
```

For local development, download the ip2region databases first:

```
sh scripts/download-data.sh
go run .
```

## Alias

Add a convenient alias to your `~/.bashrc` or `~/.zshrc`:

```bash
alias ip9='curl ip9.ir'
```

Then reload your shell config:

```bash
source ~/.bashrc  # or source ~/.zshrc
```

Now you can simply run `ip9` in your terminal to get your IP address:

```bash
ip9
```