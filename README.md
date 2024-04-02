# Concierge

_A tiny URL shortener written in Go._

> I've been exploring Go with the intention of using it to build APIs and microservices, and this is a learning project to understand Go code structure, the [Echo web framework](https://echo.labstack.com/), and [GORM](https://gorm.io) (a PostgreSQL ORM).

Concierge (_[one who keeps the entrance to an edifice](https://www.wordnik.com/words/concierge)_) is a JSON-based REST API that takes a valid alias and returns the corresponding URL via a `HTTP 301` redirect. There's also a management endpoint to create (or preview existing) records. **Once records have been created, they cannot be removed or modified via the API.**

## API Endpoints

### `/` - Root

- `GET` \
Returns documentation about the API and its listed endpoints.

### `/link` - Link Management

- `GET` - Requires a valid `alias` query-parameter. \
Returns metadata about an existing alias entry, creation date, and the corresponding URL.

- `POST` - Requires valid `alias` and `url` query-parameters. \
If the `alias` is vacant, the API will confirm creation of the new record. Returns an error message if the `alias` is occupied.

### `/to/:alias` - Link Redirection

- `GET` - Requires a valid `alias` url-parameter. \
Returns a redirection (`HTTP 301`) response to the corresponding URL via HTTPS.

## Deployment Instructions

Concierge can be run as a Docker container and uses port `3000`. \
You can pull the image [abiddiscombe/concierge](https://hub.docker.com/repository/docker/abiddiscombe/concierge/general) from Docker Hub.

The server stores data in PostgreSQL; the following environment variables are required:

- `CONCIERGE_PG_HOST` - PostgreSQL Server URL
- `CONCIERGE_PG_PORT` - PostgreSQL Server Port
- `CONCIERGE_PG_NAME` - PostgreSQL Database Name
- `CONCIERGE_PG_USER` - PostgreSQL Connection User
- `CONCIERGE_PG_PASS` - PostgreSQL Connection Password
