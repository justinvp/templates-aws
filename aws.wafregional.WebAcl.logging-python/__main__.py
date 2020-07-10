import pulumi
import pulumi_aws as aws

example = aws.wafregional.WebAcl("example", logging_configuration={
    "log_destination": aws_kinesis_firehose_delivery_stream["example"]["arn"],
    "redactedFields": {
        "fieldToMatch": [
            {
                "type": "URI",
            },
            {
                "data": "referer",
                "type": "HEADER",
            },
        ],
    },
})

