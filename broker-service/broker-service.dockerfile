FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]

# docker build -f broker-service.dockerfile -t tdboudreau/broker-service:1.0.1 . && docker push tdboudreau/broker-service:1.0.1
