FROM golang:1-alpine AS builder

WORKDIR /app

# only copy these 2 files before downloading
# docker will cache `go mod download` if they haven't changed
# resulting in faster build times, especially in larger projects with more dependencies
COPY go.mod .
COPY go.sum .
# download dependencies
RUN go mod download

# copy other files
COPY . .

# build
ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on
RUN go build -o main ./cmd/app/main.go

# move binary to a minimal image
FROM scratch
COPY --from=builder /app/main /app
COPY sql/schema /sql/schema

# run the binary
ENV PORT=80
EXPOSE 80
ENTRYPOINT ["/app"]