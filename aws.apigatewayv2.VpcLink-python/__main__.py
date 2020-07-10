import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.VpcLink("example",
    security_group_ids=[data["aws_security_group"]["example"]["id"]],
    subnet_ids=data["aws_subnet_ids"]["example"]["ids"],
    tags={
        "Usage": "example",
    })

