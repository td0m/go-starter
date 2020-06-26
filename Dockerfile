FROM golang:1.14.2-alpine AS go_builder

WORKDIR /server
COPY server /server

# build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on
RUN go build /server/cmd/server/main.go

FROM scratch 
# dont copy directly to / because static file handling would also serve linux dirs
COPY --from=go_builder /server/main /app/main
ENV PORT=80 PRODUCTION=
EXPOSE 80
ENTRYPOINT ["/app/main"]