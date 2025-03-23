FROM alpine:latest

RUN mkdir /app

COPY loggerApp /app

CMD [ "/app/loggerApp" ]

# docker build -f logger-service.dockerfile -t tdboudreau/logger-service:1.0.2 . && docker push tdboudreau/logger-service:1.0.2

# generating gRPC code from the command line
# https://grpc.io/docs/protoc-installation/
# https://grpc.io/docs/languages/go/quickstart/
# protoc --go_out=. --go_opt=paths=source_relative \
# --go-grpc_out=. --go-grpc_opt=paths=source_relative \
# routeguide/route_guide.proto