# simple-web-status-server

Extremely simple Go-based web server intended for usage by status pages. It returns a 200 OK response if a GET request to `/check_status` is sent. That's it.

## Usage

### Docker (recommended)

```bash
docker run -p 2384:2384 mrrfv/simple-web-status-server:latest

# Listen on port 8080:
docker run -p 8080:2384 mrrfv/simple-web-status-server:latest
```

### Binary

Use the `PORT` environment variable to set a different port to listen on.

```bash
go build -o ./simple-web-status-server .
./simple-web-status-server

# Listen on port 8080:
PORT=8080 ./simple-web-status-server
```

## Features

- Simple: no external dependencies or configuration required. Just run the binary or Docker container.
- Lightweight: less than 15 MB of RAM usage under high load

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Donations

If you would like to donate to support this project, you can do so via Liberapay - see my GitHub profile for the link.
