# build stage
FROM golang:1.23.2-alpine AS build

# Установите Git
RUN apk add --no-cache git

WORKDIR /src
COPY go.mod go.sum ./
# Если вы используете vendor, закомментируйте следующую строку
# RUN go mod download
COPY vendor vendor ./
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)'" -o /bin/space_vpx_satellite ./cmd/space_vpx_satellite

# prod stage
FROM alpine

# Копируем собранное приложение
COPY --from=build /bin/space_vpx_satellite /bin/space_vpx_satellite
COPY --from=build /src/internal/database/postgresql/migrations /migrations

# Укажите правильный путь к исполняемому файлу
ENTRYPOINT ["/bin/space_vpx_satellite"]