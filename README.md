# Collection of the different AWS tools written in Go

The repository contains different tools written during AWS journey with Go. The repository all the time is updated as the journey goes forward.

## Tools:

* **aws_all_users.go** - script, which gets all IAM users from your AWS account. The difference from the official documentation is explained in an article. In short: supports truncating. Don't forget to use correct ARN role with permission to list IAM users.
