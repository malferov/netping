FROM golang AS build
ARG sha
ARG ver
ARG date
WORKDIR /go/src
COPY send.go .
RUN go mod init send
RUN go build -o send -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM centos
EXPOSE 5000
COPY --from=build /go/src/send .
ENTRYPOINT ["./send"]
CMD ["-stderrthreshold=ERROR"]
