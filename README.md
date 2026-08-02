# TCP Chat

A simple Go-based chat application with separate client and server components.

## Project structure

- `client/` — terminal client that connects to the server
- `server/` — TCP server that accepts client connections

## Requirements

- Go 1.24 or newer

## Run the server

```bash
cd server
go run .
```

## Run the client

```bash
cd client
go run .
```

## Notes

- The server listens on port `3000`.
- The client prompts for a username when started.
- Both parts are currently simple prototypes and can be expanded into a full chat experience.

## License

This project is provided as-is for learning and experimentation.
