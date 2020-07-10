import pulumi
import pulumi_aws as aws

sns = aws.ses.EventDestination("sns",
    configuration_set_name=aws_ses_configuration_set["example"]["name"],
    enabled=True,
    matching_types=[
        "bounce",
        "send",
    ],
    sns_destination={
        "topic_arn": aws_sns_topic["example"]["arn"],
    })

