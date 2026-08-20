FROM node:lts-alpine AS frontend

WORKDIR /frontend-build

COPY web/package.json web/yarn.lock ./
RUN yarn install

COPY web ./
# Vite inlines VITE_* at build time, so the favicon path is a build arg.
ARG VITE_FAVICON_PATH
ENV VITE_FAVICON_PATH=$VITE_FAVICON_PATH
RUN yarn build

FROM golang:1.25-alpine AS backend

RUN apk add --no-cache gcc musl-dev linux-headers

WORKDIR /backend-build

COPY go.* ./
RUN go mod download

COPY . .
COPY --from=frontend /frontend-build/dist web/dist

RUN go build -o eth-faucet -ldflags "-s -w"

FROM alpine:3.22

# PRST-3866: patch openssl libs (libssl3/libcrypto3) for CVE-2026-45447
# (PKCS#7/S-MIME use-after-free), 3.5.6-r0 -> 3.5.7-r0. Scoped upgrade only;
# no blanket `apk upgrade`.
RUN apk add --no-cache --upgrade ca-certificates libssl3 libcrypto3

COPY --from=backend /backend-build/eth-faucet /app/eth-faucet

EXPOSE 8080

ENTRYPOINT ["/app/eth-faucet"]
