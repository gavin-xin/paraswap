FROM golang:1.17.2-alpine3.13 as build
ARG GIT_TOKEN
RUN apk add --no-cache make git
RUN echo "https://noderealbot:${GIT_TOKEN}@github.com" > ~/.git-credentials \
    && git config --global credential.helper store
# Install
ENV GIT_TERMINAL_PROMPT=1
ENV GOPRIVATE=github.com/node-real

WORKDIR /paraswap

COPY . .
RUN make build

FROM alpine:3.7

WORKDIR /opt/app/paraswap
RUN apk add --no-cache ca-certificates
COPY --from=build /paraswap/paraswap ./

ENTRYPOINT ./paraswap