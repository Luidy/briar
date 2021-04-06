FROM golang:latest
RUN mkdir /ges
ADD . /ges
WORKDIR /ges
RUN make build
ARG $PORT=10720
ENV PORT=$PORT
EXPOSE $PORT
ENTRYPOINT [ "bin/ges" ]
