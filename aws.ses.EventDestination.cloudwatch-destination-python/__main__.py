import pulumi
import pulumi_aws as aws

cloudwatch = aws.ses.EventDestination("cloudwatch",
    cloudwatch_destinations=[{
        "default_value": "default",
        "dimensionName": "dimension",
        "valueSource": "emailHeader",
    }],
    configuration_set_name=aws_ses_configuration_set["example"]["name"],
    enabled=True,
    matching_types=[
        "bounce",
        "send",
    ])

