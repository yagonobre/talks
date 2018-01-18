FROM alpine:3.6
RUN apk --no-cache --update add ca-certificates
WORKDIR /app
COPY auth /app/auth

ENTRYPOINT /app/auth
