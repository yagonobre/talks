FROM golang
RUN go get github.com/aws/aws-sdk-go/aws
ADD main.go .
RUN go build main.go
CMD "./main"
