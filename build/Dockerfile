FROM alpine
RUN apk add --no-cache ca-certificates
ADD auth-service .
CMD ["./auth-service"]