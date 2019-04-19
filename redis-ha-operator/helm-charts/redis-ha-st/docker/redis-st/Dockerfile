FROM registry.access.redhat.com/ubi7-dev-preview/ubi-minimal:7.6

ENV BIN_FILE=/usr/local/bin/sentinel_tunnel \
    ETC_FILE=/etc/sentinel_tunnel/sentinel_tunnel.json

COPY sentinel_tunnel ${BIN_FILE}
RUN mkdir -p /etc/sentinel_tunnel
COPY sentinel_tunnel.json ${ETC_FILE}

VOLUME ["/etc/sentinel_tunnel"]

ENTRYPOINT ["/usr/local/bin/sentinel_tunnel", "/etc/sentinel_tunnel/sentinel_tunnel.json", "/dev/stdout"]