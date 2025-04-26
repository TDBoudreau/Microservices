FROM caddy:2.4.6-alpine

COPY Caddyfile /etc/caddy/Caddyfile

# docker build -f caddy.dockerfile -t tdboudreau/micro-caddy:1.0.0 . && docker push tdboudreau/micro-caddy:1.0.0

# Cross platform linux/amd64,linux/arm64
# docker buildx build -f caddy.dockerfile --platform linux/amd64,linux/arm64 -t tdboudreau/micro-caddy:1.0.1 --push .