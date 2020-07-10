import pulumi
import pulumi_aws as aws

kinesis = aws.ses.EventDestination("kinesis",
    configuration_set_name=aws_ses_configuration_set["example"]["name"],
    enabled=True,
    kinesis_destination={
        "role_arn": aws_iam_role["example"]["arn"],
        "stream_arn": aws_kinesis_firehose_delivery_stream["example"]["arn"],
    },
    matching_types=[
        "bounce",
        "send",
    ])

