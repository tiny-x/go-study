FROM ubuntu:latest
LABEL maintainer="Fei ye"

# install gcc make git
RUN apt-get update \
    && apt-get install -y build-essential git xz-utils \
    && git clone git://github.com/kongjian/tsar.git \
    && cd tsar \
    && make \
    && make install

CMD [ "top" ]
