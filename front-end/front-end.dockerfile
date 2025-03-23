FROM alpine:latest

RUN mkdir /app

COPY frontApp /app

CMD [ "/app/frontApp" ]

# docker build -f front-end.dockerfile -t tdboudreau/front-end:1.0.2 . && docker push tdboudreau/front-end:1.0.2
