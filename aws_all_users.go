// LICENSE: MIT
// Author: Viktorija Almazova

// Script gets all AWS IAM users, taking into account truncating. Just to be sure that we haven't missed anyone :)
// It supposed to run with ARN applied taken from an environment variable. ARN role shall have permission to list Users (iam:ListUsers)
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

// AWS role needed to check IAM
var awsArnRole = os.Getenv("ARNROLE")

func main() {
	// Apply role and read all IAM users
	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, awsArnRole)

	svc := iam.New(sess, &aws.Config{Credentials: creds})

	var marker *string
	collectedUsers := []*iam.ListUsersOutput{}

	for {

		response, err := svc.ListUsers(&iam.ListUsersInput{Marker: marker})

		if err != nil {
			fmt.Print(err)
		}

		collectedUsers = append(collectedUsers, response)

		if *response.IsTruncated == false {
			break
		}

		marker = response.Marker

	}

	for _, userSet := range collectedUsers {
		for _, user := range userSet.Users {
			fmt.Printf("AWS IAM username: %s\n", *user.UserName)
		}
	}

}
