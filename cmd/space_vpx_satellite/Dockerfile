FROM golang:alpine AS build


RUN apk add --no-cache git

WORKDIR /src
COPY go.mod go.sum ./
COPY vendor ./vendor
COPY ./ ./




RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)'" -o /bin/space_vpx_satellite ./cmd/space_vpx_satellite

FROM alpine

COPY --from=build /bin/space_vpx_satellite /bin/space_vpx_satellite
COPY --from=build /src/internal/database/postgresql/migrations /migrations

ENTRYPOINT ["/bin/space_vpx_satellite"]