FROM golang:1.12-alpine
WORKDIR /root

RUN apk add git make

RUN mkdir -p ./temp/ag7if/
COPY ./ ./temp/ag7if/
RUN cd ./temp/ag7if/ && make clean
RUN cd ./temp/ag7if/ && make
RUN cd ./temp/ag7if/ && make install
RUN rm -rf ./temp/

CMD ["ag7if"]