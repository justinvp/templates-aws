import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafregional.WebAcl("example", {
    loggingConfiguration: {
        logDestination: aws_kinesis_firehose_delivery_stream_example.arn,
        redactedFields: {
            fieldToMatches: [
                {
                    type: "URI",
                },
                {
                    data: "referer",
                    type: "HEADER",
                },
            ],
        },
    },
});

