FROM golang AS build
ARG sha
ARG ver
ARG date
ARG CGO_ENABLED=0
WORKDIR /go/src
COPY send.go .
RUN go mod init send
RUN go mod tidy
RUN go build -o send -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM centos
EXPOSE 5000
COPY --from=build /go/src/send .
ENTRYPOINT ["./send"]
CMD ["-stderrthreshold=ERROR"]
