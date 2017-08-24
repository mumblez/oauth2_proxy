FROM golang:1.8.3-alpine
COPY oauth2_proxy /oauth2_proxy
CMD mkdir /etc/oauth
# mount secrets into /etc/oauth/{oauth2_proxy.cfg,inf-oauth.json}
EXPOSE 4128
ENTRYPOINT ["/oauth2_proxy", "-config", "/etc/oauth/oauth2_proxy.cfg"]
