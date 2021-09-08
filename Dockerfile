FROM golang:1.14.9-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum main.go mongo.go recipes.json /build/
WORKDIR /build
RUN go build -o recipeapi

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/recipeapi /app/
COPY --from=builder /build/recipes.json /app/
WORKDIR /app
CMD ["./recipeapi"]