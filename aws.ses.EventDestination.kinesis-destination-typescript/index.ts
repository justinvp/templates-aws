import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const kinesis = new aws.ses.EventDestination("kinesis", {
    configurationSetName: aws_ses_configuration_set_example.name,
    enabled: true,
    kinesisDestination: {
        roleArn: aws_iam_role_example.arn,
        streamArn: aws_kinesis_firehose_delivery_stream_example.arn,
    },
    matchingTypes: [
        "bounce",
        "send",
    ],
});

