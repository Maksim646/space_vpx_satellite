FROM golang:alpine AS build


RUN apk add --no-cache git

WORKDIR /src
COPY space_vpx_satellite/go.mod space_vpx_satellite/go.sum ./
COPY space_vpx_satellite/vendor ./vendor
COPY space_vpx_satellite/ ./


RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)'" -o /bin/space_vpx_satellite ./cmd/space_vpx_satellite

FROM alpine

COPY --from=build /bin/space_vpx_satellite /bin/space_vpx_satellite
COPY --from=build /src/internal/database/postgresql/migrations /migrations

ENTRYPOINT ["/bin/space_vpx_satellite"]