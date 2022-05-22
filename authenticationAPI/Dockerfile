FROM golang:1.18.2-alpine3.15 as build
RUN mkdir /app
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=build /app/app ./
CMD ["./app"]