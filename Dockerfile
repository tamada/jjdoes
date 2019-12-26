FROM alpine:3.10.1
LABEL maintainer="Haruaki Tamada" \
      tjdoe-version="1.0.0" \
      description=""

RUN    adduser -D tjdoe \
    && apk --no-cache add curl=7.66.0-r0 tar=1.32-r0 \
    && curl -s -L -O https://github.com/tamada/tjdoe/releases/download/v1.0.0/tjdoe-1.0.0_linux_amd64.tar.gz \
    && tar xfz tjdoe-1.0.0_linux_amd64.tar.gz  \
    && mv tjdoe-1.0.0 /opt                     \
    && ln -s /opt/tjdoe-1.0.0 /opt/tjdoe       \
    && ln -s /opt/tjdoe /usr/local/share/tjdoe \
    && rm tjdoe-1.0.0_linux_amd64.tar.gz       \
    && ln -s /opt/tjdoe/tjdoe /usr/local/bin/tjdoe

ENV HOME="/home/tjdoe"

WORKDIR /home/tjdoe
USER    tjdoe

ENTRYPOINT [ "tjdoe" ]
