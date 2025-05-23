FROM alpine:latest

RUN mkdir /app

COPY mailerApp /app

CMD [ "/app/mailerApp" ]

# docker build -f mail-service.dockerfile -t tdboudreau/mail-service:1.0.2 --push .

# Cross platform linux/amd64,linux/arm64
# docker buildx build -f mail-service.dockerfile --platform linux/amd64,linux/arm64 -t tdboudreau/mail-service:1.0.2 --push .