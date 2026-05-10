# AINews

`ainews.personastack.ai` is a small Go web service that publishes current AI news briefs.

## Endpoints

- `/` renders the current homepage and article links.
- `/posts/{slug}` renders a full article page.
- `/api/posts` returns the published post list as JSON.
- `/healthz` returns `ok`.

## Local Run

Set `PORT` to override the default `8080`, then run `./cmd/ainews`.

## Build

The repository includes a multi-stage `Dockerfile` that runs `go test ./...` before building the final binary image.
