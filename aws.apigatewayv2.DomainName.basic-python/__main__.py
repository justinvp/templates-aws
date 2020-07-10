import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.DomainName("example",
    domain_name="ws-api.example.com",
    domain_name_configuration={
        "certificate_arn": aws_acm_certificate["example"]["arn"],
        "endpoint_type": "REGIONAL",
        "security_policy": "TLS_1_2",
    })

