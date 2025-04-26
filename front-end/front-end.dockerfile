FROM alpine:latest

RUN mkdir /app

COPY frontApp /app

CMD [ "/app/frontApp" ]

# Cross platform linux/amd64,linux/arm64
# docker build -f front-end.dockerfile -t tdboudreau/front-end:1.0.4 --push .
# docker buildx build -f front-end.dockerfile --platform linux/amd64,linux/arm64 -t tdboudreau/front-end:1.0.4 --push .