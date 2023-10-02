FROM alpine:latest

EXPOSE 8000

# The command to run
CMD ["/devs-crud"]

ARG BUILD_TAG=unknown
LABEL BUILD_TAG=$BUILD_TAG

COPY bin/devs-crud /devs-crud

