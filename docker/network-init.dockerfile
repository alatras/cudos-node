FROM golang:alpine

RUN apk add --no-cache make bash

WORKDIR /usr/cudos

COPY ./project-node ./

RUN make

RUN chmod +x ./init-root.sh

CMD ["/bin/sh", "-c", "./init-network.sh&& cudos-noded start"] 