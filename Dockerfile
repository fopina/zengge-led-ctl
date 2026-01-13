FROM alpine:3.23

COPY zengge-led-ctl /usr/bin/zengge-led-ctl
ENTRYPOINT ["/usr/bin/zengge-led-ctl"]