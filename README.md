# Go-Starter

An unopinionated Go project template with useful utilities, good file structure, docker integration, and more.

Startup, side project, or hackathon starter pack. For those familiar with Go, Docker, SQL, and JWT authentication.

## Usage

### Development

The easiest way to run the app is via the `docker-compose.yml` config.

This command will rebuild and start all services:

```bash
docker-compose up --build
```

However, if for some reason you'd like to set up the services yourself outside of docker, you can run the application using the following command:

```bash
go run ./cmd/app/main.go
```

Keep in mind though that you will still need to start and configure any required services such as `PostgresQL`.

You can either start them locally or within docker:

```bash
docker-compose up -d postgres
```

### Production

Docker Compose is well suited for production, as long as:
 - you're deploying to 1 host
 - you don't mind a few seconds of downtime each time you deploy a new version (no load balancer)
 - you don't need vertical scaling

If you need any of the above, you should probably look into using Kubernetes.
