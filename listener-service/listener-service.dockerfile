FROM alpine:latest

RUN mkdir /app

COPY listenerApp /app

CMD [ "/app/listenerApp" ]

# docker build -f listener-service.dockerfile -t tdboudreau/listener-service:1.0.1 --push .

# Cross platform linux/amd64,linux/arm64
# docker buildx build -f listener-service.dockerfile --platform linux/amd64,linux/arm64 -t tdboudreau/listener-service:1.0.1 --push .