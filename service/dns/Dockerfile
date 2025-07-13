FROM golang AS build
ARG sha
ARG ver
ARG date
ARG CGO_ENABLED=0
WORKDIR /go/src
COPY dns.go .
RUN go mod init dns
RUN go mod tidy
RUN go build -o dns -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM alpine
RUN apk add bind-tools
EXPOSE 5000
COPY --from=build /go/src/dns .
COPY dig.sh .
ENTRYPOINT ["./dns"]
CMD ["-stderrthreshold=ERROR"]
