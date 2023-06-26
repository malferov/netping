FROM golang AS build
ARG sha
ARG ver
ARG date
ARG CGO_ENABLED=0
WORKDIR /go/src
COPY uuid.go .
RUN go mod init uuid
RUN go mod tidy
RUN go build -o uuid -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM centos
EXPOSE 5000
COPY --from=build /go/src/uuid .
ENTRYPOINT ["./uuid"]
CMD ["-stderrthreshold=ERROR"]
