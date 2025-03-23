FROM alpine:latest

RUN mkdir /app

COPY mailerApp /app

CMD [ "/app/mailerApp" ]

# docker build -f mail-service.dockerfile -t tdboudreau/mail-service:1.0.1 . && docker push tdboudreau/mail-service:1.0.1
