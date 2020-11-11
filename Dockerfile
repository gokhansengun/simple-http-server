FROM golang:1.15.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/server .

FROM alpine:3.12.1 AS runtime
COPY --from=build /out/server /

CMD [ "/server" ]
