FROM golang:1.17 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

ARG version
RUN go build -ldflags="-X github.com/jaunty/jaunty/internal/pkg/version.version=${version}" -o jaunty

FROM gcr.io/distroless/base:nonroot
COPY --from=builder /app/jaunty /jaunty
ENTRYPOINT [ "/jaunty" ]