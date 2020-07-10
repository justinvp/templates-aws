import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cloudwatchRole = new aws.iam.Role("cloudwatch", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "apigateway.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
`,
});
const demo = new aws.apigateway.Account("demo", {
    cloudwatchRoleArn: cloudwatchRole.arn,
});
const cloudwatchRolePolicy = new aws.iam.RolePolicy("cloudwatch", {
    policy: `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:DescribeLogGroups",
                "logs:DescribeLogStreams",
                "logs:PutLogEvents",
                "logs:GetLogEvents",
                "logs:FilterLogEvents"
            ],
            "Resource": "*"
        }
    ]
}
`,
    role: cloudwatchRole.id,
});

