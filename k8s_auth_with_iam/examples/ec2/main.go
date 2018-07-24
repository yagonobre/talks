package main

import (
    "fmt"
	"time"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
    // Load session from shared config
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create new EC2 client
    ec2Svc := ec2.New(sess)

	for {
		result, err := ec2Svc.DescribeInstances(nil)
		 if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("Success", result)
		}
		time.Sleep(1000 * time.Millisecond)
	}

}
