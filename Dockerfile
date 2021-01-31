FROM golang:1.15-alpine as build
RUN apk --no-cache add git
WORKDIR /app
COPY . /app
RUN go build -o ./zamio ./cmd/zamio/main.go

FROM alpine:3.10.3
WORKDIR /app
COPY --from=build /app/.env ./
COPY --from=build /app/zamio ./
RUN chmod +x ./zamio
CMD ["./zamio"]