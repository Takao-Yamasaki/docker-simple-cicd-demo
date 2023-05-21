# local
FROM golang:1.20 AS base
CMD ["bash"]

# test-and-build
FROM base AS test-and-build
COPY . /workspace
WORKDIR /workspace
RUN go test -v && go build -o main .

# production
FROM ubuntu:latest AS production
COPY --from=test-and-build /workspace/main /app/main
WORKDIR /app

CMD ["./main"]
