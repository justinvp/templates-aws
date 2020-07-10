import pulumi
import pulumi_aws as aws

cert = aws.iot.Certificate("cert", active=True)

