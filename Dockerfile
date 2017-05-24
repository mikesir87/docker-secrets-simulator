FROM golang AS build
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

FROM scratch
COPY --from=build /app/main /main
CMD ["/main"]
