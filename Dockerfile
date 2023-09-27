FROM golang:alpine AS personal-builder
WORKDIR /app
RUN apk add --no-cache hugo
COPY personal/ .
RUN hugo

FROM node:lts-alpine AS staffing-builder
WORKDIR /app
COPY staffing/ .
RUN npm install
ENV NEXT_TELEMETRY_DISABLED=1
RUN npx next build
RUN npx next export

FROM golang:alpine AS main-builder
WORKDIR /app
COPY *.go go.* ./
COPY --from=personal-builder /app/public ./personal/public
COPY --from=staffing-builder /app/out ./staffing/out
RUN go build ./

FROM alpine:latest
WORKDIR /app
COPY --from=main-builder /app/Website .

ENTRYPOINT ["./Website"]

EXPOSE 8080