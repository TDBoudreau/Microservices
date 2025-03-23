FROM alpine:latest

RUN mkdir /app

COPY listenerApp /app

CMD [ "/app/listenerApp" ]

# docker build -f listener-service.dockerfile -t tdboudreau/listener-service:1.0.0 . && docker push tdboudreau/listener-service:1.0.0
