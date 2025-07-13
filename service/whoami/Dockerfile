FROM golang AS build
ARG sha
ARG ver
ARG date
ARG CGO_ENABLED=0
WORKDIR /go/src
COPY whoami.go .
RUN go mod init whoami
RUN go mod tidy
RUN go build -o whoami -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM alpine
EXPOSE 5000
COPY --from=build /go/src/whoami .
ENTRYPOINT ["./whoami"]
CMD ["-stderrthreshold=ERROR"]
