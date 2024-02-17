FROM golang:1.21.5 as builder

WORKDIR /build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags "-s -w" -o simple-responder

# -----------------------------------------------------------------------------
FROM alpine as runtime

ENV GIN_MODE=release

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=builder /build/simple-responder /usr/local/sbin/simple-responder

CMD ["simple-responder"]
