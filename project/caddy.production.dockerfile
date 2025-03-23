FROM caddy:2.4.6-alpine

COPY Caddyfile.production /etc/caddy/Caddyfile

# docker build -f caddy.production.dockerfile -t tdboudreau/micro-caddy-production:1.0.1 . && docker push tdboudreau/micro-caddy-production:1.0.1