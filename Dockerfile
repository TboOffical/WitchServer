FROM golang:latest
RUN git clone https://github.com/TboOffical/WitchServer.git
RUN sed -i 's/_host := "localhost"/_host := "0.0.0.0"/g' /go/WitchServer/witch.go
RUN cd WitchServer && go mod init Witch && go get -u github.com/gen2brain/dlgs && go build ./witch.go ./util.go ./wba.go ./listener.go
EXPOSE 8000
ENTRYPOINT [ "/go/WitchServer/witch", "8000" ]
