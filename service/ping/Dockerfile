FROM golang:1.15 AS build
ARG sha
ARG ver
ARG date
WORKDIR /go/src
COPY ping.go .
RUN go mod init ping
RUN go mod tidy
RUN go build -o ping -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM centos
EXPOSE 5000
COPY --from=build /go/src/ping .
ENTRYPOINT ["./ping"]
CMD ["-stderrthreshold=ERROR"]
