package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"

    "fmt"
    "os"
	"time"
)

func main() {
    ids := []string{"sg-089dcdb7ab7736702"}

	sess, _ := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1")},
    )

    // Create an EC2 service client.
    svc := ec2.New(sess)

	for {
		result, err := svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
			GroupIds: aws.StringSlice(ids),
		})
		if err != nil {
			fmt.Println("Error", err)
		}

		fmt.Println("Security Group:")
		for _, group := range result.SecurityGroups {
			fmt.Println(group)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}
