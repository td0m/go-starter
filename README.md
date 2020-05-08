# Go Starter

Scalable, non-opinionated backend starter.

## Features
 - Docker
 - PostgresQL

## Architecture

This template follows Domain Driven Development (DDD) Hegaxonal Architecture. If you've never heard of it, you should definitely [watch this talk from GopherCon](https://www.youtube.com/watch?v=oL6JBUk6tj0).

The aim is to separate the **logic** in the app, from the **input** and **output** interfaces. This allows not only maintainability as the project scales, but it also allows you to extend the application (by adding for example `grpc` in addition to `rest` api endpoint) or swap the database provider (e.g. from SQL to Postgres) without having to worry about the logic.

Structure is grouped by context, and each time the interfaces used are declared.