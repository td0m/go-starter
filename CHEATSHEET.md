# Command Cheatsheet

## Docker Compose

| Command                                   | Description                                                    |
| ----------------------------------------- | -------------------------------------------------------------- |
| `docker-compose ps`                       | List containers, their state, exposed ports, etc.              |
| `docker-compose up`                       | Start all containers in attached mode.                         |
| `docker-compose up C1 C2 ...`             | Start containers in attached mode.                             |
| `docker-compose up -d`                    | Start containers in detached mode.                             |
| `docker-compose logs`                     | View logs of all detached containers.                          |
| `docker-compose logs C1 C2 ...`           | View logs of only specified containers.                        |
| `docker-compose logs --tail=1000`         | View last 1000 lines of logs.                                  |
| `docker-compose logs --follow`            | Shows new logs instead of quitting                             |
| `docker-compose stop`                     | Stop all containers.                                           |
| `docker-compose pull C1`                  | Pull container.                                                |
| `docker-compose down --rmi all --volumes` | Stop and remove all containers, networks, images, and volumes. |
| `docker-compose -f path/to/file up`       | Use a different config file to run `up` command                |

### Override

`docker-compose.override.yml` will automatically merge with the default compose file unless the `-f` flag is used.

This can be useful for development-only environments, for example to attach volumes or bind to host's ports.

