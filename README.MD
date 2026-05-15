# PokédexGo

A CLI Pokédex REPL written in Go that hits the [PokéAPI](https://pokeapi.co).

## Commands

| Command | Description |
|---|---|
| `map` / `mapb` | page forward/backward through location areas |
| `explore <location>` | list all Pokémon in a location area |
| `catch <pokemon>` | attempt to catch a Pokémon — catch rate is based on `base_experience`, so rarer Pokémon are harder to catch |
| `inspect <pokemon>` | print stats, types, height, and weight of a caught Pokémon |
| `pokedex` | list all caught Pokémon |

## Usage

```bash
go run .
```

## How it works

API responses are cached in-memory using a `map[string]cacheEntry` protected by a `sync.Mutex`. A background goroutine runs a ticker that evicts entries older than the configured interval, so repeated commands don't hit the network unnecessarily.

## What I learned

- Building a REPL loop in Go — reading stdin, parsing input, dispatching to command handlers
- Structuring a Go project with internal packages (`pokeapi`, `pokecache`)
- Writing a thread-safe TTL cache from scratch using `sync.Mutex` and a background cleanup goroutine
- Consuming a REST API in Go — `http.Client`, `json.Unmarshal`, and handling paginated responses with next/previous URLs
- Go interfaces, structs, and pointer receivers
