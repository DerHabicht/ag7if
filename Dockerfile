FROM golang:1.14-alpine
WORKDIR /root

RUN apk add git make

RUN mkdir -p ./temp/ag7if/
COPY ./ ./temp/ag7if/
RUN cd ./temp/ag7if/ && make clean
RUN cd ./temp/ag7if/ && make
RUN cd ./temp/ag7if/ && make install
RUN rm -rf ./temp/
RUN rm -rf /go/src/

EXPOSE 8080
ENV URL 0.0.0.0:8080

CMD ["/go/bin/ag7if"]
