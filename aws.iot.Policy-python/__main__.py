import pulumi
import pulumi_aws as aws

pubsub = aws.iot.Policy("pubsub", policy="""{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "iot:*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}

""")

