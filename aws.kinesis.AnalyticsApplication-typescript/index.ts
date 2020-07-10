import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testStream = new aws.kinesis.Stream("test_stream", {
    shardCount: 1,
});
const testApplication = new aws.kinesis.AnalyticsApplication("test_application", {
    inputs: {
        kinesisStream: {
            resourceArn: testStream.arn,
            roleArn: aws_iam_role_test.arn,
        },
        namePrefix: "test_prefix",
        parallelism: {
            count: 1,
        },
        schema: {
            recordColumns: [{
                mapping: "$.test",
                name: "test",
                sqlType: "VARCHAR(8)",
            }],
            recordEncoding: "UTF-8",
            recordFormat: {
                mappingParameters: {
                    json: {
                        recordRowPath: "$",
                    },
                },
            },
        },
    },
});

