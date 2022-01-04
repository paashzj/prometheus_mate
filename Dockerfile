#
# Build stage
#
FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o prom_mate .


FROM ttbb/prometheus:nake

LABEL maintainer="shoothzj@gmail.com"

COPY pkg/config/gf.toml /opt/sh/prometheus/mate/config/gf.toml

COPY docker-build /opt/sh/prometheus/mate

COPY --from=build /opt/sh/compile/pkg/prom_mate /opt/sh/prometheus/mate/prom_mate

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/prometheus/mate/scripts/start.sh"]