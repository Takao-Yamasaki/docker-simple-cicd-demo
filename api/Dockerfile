FROM golang:1.20 AS base
CMD ["bash"]

# build
FROM base AS build
COPY . /workspace
WORKDIR /workspace
RUN go build main.go

# production
FROM alpine:latest AS production
COPY --from=build /workspace/main .
CMD ["./main"]
