# local
FROM golang:1.20-alpine3.18 AS base
CMD [ "bash" ]

# test-and-build
FROM base AS test-and-build
COPY . /workspace
WORKDIR /workspace
RUN go mod download

# CGO_ENABLED=0: CGOを無効化する。静的バイナリを作成する際は無効化が推奨。
# GOOS=linux: ターゲットOSをLinuxにする
# -a -installsuffix: ビルドされたバイナリにCGOの情報を含める
# RUN go test -v && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .
RUN go test -v && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# production
FROM alpine:3.18 AS production
COPY --from=test-and-build /workspace/main /app/main
WORKDIR /app
CMD [ "./main" ]
