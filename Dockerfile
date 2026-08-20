FROM node:lts-alpine AS frontend

# Brand selected at build time: docker build --build-arg VITE_BRAND=teiza
ARG VITE_BRAND=base

WORKDIR /frontend-build

COPY web/package.json web/yarn.lock ./
RUN yarn install

COPY web ./
RUN VITE_BRAND="$VITE_BRAND" yarn build

FROM golang:1.25-alpine AS backend

RUN apk add --no-cache gcc musl-dev linux-headers

WORKDIR /backend-build

COPY go.* ./
RUN go mod download

COPY . .
COPY --from=frontend /frontend-build/dist web/dist

RUN go build -o eth-faucet -ldflags "-s -w"

FROM alpine:3.22

RUN apk add --no-cache ca-certificates

COPY --from=backend /backend-build/eth-faucet /app/eth-faucet

EXPOSE 8080

ENTRYPOINT ["/app/eth-faucet"]
