import pulumi
import pulumi_aws as aws

demo = aws.apigateway.ClientCertificate("demo", description="My client certificate")

