FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]

# Cross platform linux/amd64,linux/arm64
# docker build -f broker-service.dockerfile -t tdboudreau/broker-service:1.0.3 --push .
# docker buildx build -f broker-service.dockerfile --platform linux/amd64,linux/arm64 -t tdboudreau/broker-service:1.0.3 --push .