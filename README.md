# Concierge

_A tiny URL shortener written in Go._

> I've been exploring Go with the intention of using it to build microservices; this is a learning project to understand Go code structure, the [Echo web framework](https://echo.labstack.com/), and [GORM](https://gorm.io) (a PostgreSQL ORM).

Concierge (_"[one who keeps the entrance to an edifice](https://www.wordnik.com/words/concierge)"_) is a JSON-based REST API which takes a valid alias and returns the corresponding URL via a `HTTP 301` redirect. There's also a management endpoint to create (or preview existing) records. **Once records are created, they can't removed/modified via the API.**

## API Specification

### `/` - Root

- `GET` \
Returns documentation about the API and its endpoints.

### `/link/:alias` - Link Management

- `GET` - Requires an `alias` URL parameter. \
Returns metadata about an existing entry matching the alias, its creation date, and the corresponding URL.

- `POST` - Requires an `alias` URL parameter and a `url` query parameter. \
If the `alias` is vacant, a new unique record will be created for the provided `url`.

### `/to/:alias` - Link Redirection

- `GET` - Requires an `alias` URL parameter. \
**⭐️ Share links to this endpoint with your target audience.⭐️** \
Returns a redirection `HTTP 301` response to the corresponding URL of the supplied alias. All redirections are prefixed with `https://` automatically.

## Deployment Instructions

Concierge can be run as a Docker container and uses internal port `:3000`. \
The latest release image of [abiddiscombe/concierge](https://hub.docker.com/repository/docker/abiddiscombe/concierge/general) can be pulled from Docker Hub.

The API server is backed by PostgreSQL; the following environment variables are required:

- `CONCIERGE_PG_HOST` - PostgreSQL Server URL
- `CONCIERGE_PG_PORT` - PostgreSQL Server Port
- `CONCIERGE_PG_NAME` - PostgreSQL Database Name
- `CONCIERGE_PG_USER` - PostgreSQL Connection User
- `CONCIERGE_PG_PASS` - PostgreSQL Connection Password
