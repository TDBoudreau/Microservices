FROM caddy:2.4.6-alpine

COPY Caddyfile.production /etc/caddy/Caddyfile

# Old images
# docker build -f caddy.production.dockerfile -t tdboudreau/micro-caddy-production:1.0.1 . && docker push tdboudreau/micro-caddy-production:1.0.1

# Cross platform linux/amd64,linux/arm64
# docker buildx build -f caddy.production.dockerfile --platform linux/amd64,linux/arm64 -t tdboudreau/micro-caddy-production:1.0.2 --push .