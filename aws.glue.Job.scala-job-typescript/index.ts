import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Job("example", {
    command: {
        scriptLocation: pulumi.interpolate`s3://${aws_s3_bucket_example.bucket}/example.scala`,
    },
    defaultArguments: {
        "--job-language": "scala",
    },
    roleArn: aws_iam_role_example.arn,
});

