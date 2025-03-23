FROM caddy:2.4.6-alpine

COPY Caddyfile /etc/caddy/Caddyfile

# docker build -f caddy.dockerfile -t tdboudreau/micro-caddy:1.0.0 . && docker push tdboudreau/micro-caddy:1.0.0