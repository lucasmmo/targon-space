FROM golang:1.16 as builder
WORKDIR /build
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app cmd/rest/main.go


FROM alpine:latest
EXPOSE 8080
WORKDIR /rest
COPY --from=builder /build/.env .
COPY --from=builder /build/app .
CMD [ "./app" ]