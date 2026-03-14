# IP9
A death simple Golang program, that returns your IP address.

Build and run with docker-compose:

```
docker compose up -d
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