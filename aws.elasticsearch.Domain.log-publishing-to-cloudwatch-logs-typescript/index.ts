import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("example", {});
const exampleLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("example", {
    policyDocument: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "es.amazonaws.com"
      },
      "Action": [
        "logs:PutLogEvents",
        "logs:PutLogEventsBatch",
        "logs:CreateLogStream"
      ],
      "Resource": "arn:aws:logs:*"
    }
  ]
}
`,
    policyName: "example",
});
const exampleDomain = new aws.elasticsearch.Domain("example", {
    logPublishingOptions: [{
        cloudwatchLogGroupArn: exampleLogGroup.arn,
        logType: "INDEX_SLOW_LOGS",
    }],
});

