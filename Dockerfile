FROM golang:latest
RUN mkdir /ges
ADD . /ges
WORKDIR /ges
RUN make build
ARG $PORT=10630
ENV PORT=$PORT
EXPOSE $PORT
ENTRYPOINT [ "bin/ges" ]
