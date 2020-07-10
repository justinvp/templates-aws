import pulumi
import pulumi_aws as aws

example = aws.ec2clientvpn.Endpoint("example",
    authentication_options=[{
        "rootCertificateChainArn": aws_acm_certificate["root_cert"]["arn"],
        "type": "certificate-authentication",
    }],
    client_cidr_block="10.0.0.0/16",
    connection_log_options={
        "cloudwatchLogGroup": aws_cloudwatch_log_group["lg"]["name"],
        "cloudwatchLogStream": aws_cloudwatch_log_stream["ls"]["name"],
        "enabled": True,
    },
    description="clientvpn-example",
    server_certificate_arn=aws_acm_certificate["cert"]["arn"])

