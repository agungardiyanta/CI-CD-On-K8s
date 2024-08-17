FROM golang:1.22-alpine as build
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o binary

FROM golang:1.22-alpine as final
WORKDIR /app
COPY --from=build /app/binary/ /app/
COPY --from=build /app/templates /app/templates
EXPOSE 8080
ENTRYPOINT  [ "./binary" ]