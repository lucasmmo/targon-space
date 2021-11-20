FROM golang:1.16 as builder
WORKDIR /build
COPY .  . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app cmd/main.go


FROM alpine:latest
COPY --from=builder /build/app .
CMD [ "./app" ]