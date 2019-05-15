FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./ /go/src/github.com/lucas-clemente/quic-go
WORKDIR /go/src/github.com/lucas-clemente/quic-go
RUN git remote add mp-quic https://github.com/MatheusDonizete/mp-quic
RUN git fetch mp-quic
RUN git checkout sc-eval
RUN go get -t -u ./...

RUN go build