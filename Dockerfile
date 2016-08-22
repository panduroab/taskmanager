#golang image. Workspace GOPATH configured at /go
FROM golang

#Copy local package files to the container's workspace
ADD . /go/src/github.com/panduroab/taskmanager

#Setting up workspace directory
WORKDIR /go/src/github.com/panduroab/taskmanager

#Get godeps for managing and restoring dependencies
RUN go get github.com/tools/godep

#Restore godep dependencies
RUN godep restore

#Build the taskmanager command
RUN go install github.com/panduroab/taskmanager

#Run the taskmanager command when the container starts
ENTRYPOINT /go/bin/taskmanager

#Service listen on port 8080
EXPOSE 8080
