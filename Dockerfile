FROM alpine:3.22

COPY zengge-led-ctl /usr/bin/zengge-led-ctl
ENTRYPOINT ["/usr/bin/zengge-led-ctl"]