FROM alpine:3.10.1

ARG version="1.0.0"

LABEL maintainer="Haruaki Tamada" \
      tjdoe-version="${version}" \
      description="anonymizing given programs for programming courses and their scores for grades."

RUN    adduser -D tjdoe \
    && apk update       \
    && apk --no-cache add --virtual .builddeps curl tar \
    && curl -s -L -O https://github.com/tamada/tjdoe/releases/download/v${version}/tjdoe-${version}_linux_amd64.tar.gz \
    && tar xfz tjdoe-${version}_linux_amd64.tar.gz \
    && mv tjdoe-${version} /opt                    \
    && ln -s /opt/tjdoe-${version} /opt/tjdoe      \
    && rm tjdoe-${version}_linux_amd64.tar.gz      \
    && ln -s /opt/tjdoe/tjdoe /usr/local/bin/tjdoe \
    && apk del --purge .builddeps

ENV HOME="/home/tjdoe"

WORKDIR /home/tjdoe
USER    tjdoe

ENTRYPOINT [ "tjdoe" ]
