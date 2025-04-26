FROM alpine:latest

RUN mkdir /app

COPY authApp /app

CMD [ "/app/authApp" ]

# Cross platform linux/amd64,linux/arm64
# docker build -f authentication-service.dockerfile -t tdboudreau/authentication-service:1.0.3 --push .
# docker buildx build -f authentication-service.dockerfile --platform linux/amd64,linux/arm64 -t tdboudreau/authentication-service:1.0.3 --push .