FROM quay.io/mirrorhub/devcontainer

USER root
WORKDIR /root
RUN git clone https://github.com/mirrorhub-io/platform
RUN mkdir -p /go/src/github.com/mirrorhub-io
RUN ln -s /root/platform /go/src/github.com/mirrorhub-io/platform
WORKDIR /root/platform
RUN go build .

EXPOSE 9000 8080
CMD ./platform api
