FROM golang:1.16 as builder
WORKDIR /build
ADD * /build/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .


FROM alpine:latest
COPY --from=builder /build/app .
CMD [ "./app" ]