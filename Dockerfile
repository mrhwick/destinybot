FROM alpine:latest

VOLUME /usr/config

WORKDIR /usr/src

ADD destinybot .

ENTRYPOINT ["./destinybot"]

CMD ["--config", "/usr/config/config.yaml"]


