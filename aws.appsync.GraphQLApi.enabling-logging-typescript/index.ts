import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRole = new aws.iam.Role("example", {
    assumeRolePolicy: `{
    "Version": "2012-10-17",
    "Statement": [
        {
        "Effect": "Allow",
        "Principal": {
            "Service": "appsync.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
        }
    ]
}
`,
});
const exampleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("example", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSAppSyncPushToCloudWatchLogs",
    role: exampleRole.name,
});
const exampleGraphQLApi = new aws.appsync.GraphQLApi("example", {
    logConfig: {
        cloudwatchLogsRoleArn: exampleRole.arn,
        fieldLogLevel: "ERROR",
    },
});

