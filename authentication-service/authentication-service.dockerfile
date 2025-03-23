FROM alpine:latest

RUN mkdir /app

COPY authApp /app

CMD [ "/app/authApp" ]

# docker build -f authentication-service.dockerfile -t tdboudreau/authentication-service:1.0.1 . && docker push tdboudreau/authentication-service:1.0.1
