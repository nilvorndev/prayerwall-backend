FROM golang:1.27-alpine AS build
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/prayerwall-api ./cmd/api

FROM alpine:3.22
RUN apk add --no-cache ca-certificates
COPY --from=build /out/prayerwall-api /usr/local/bin/prayerwall-api
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/prayerwall-api"]
