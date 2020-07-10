import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.DomainName("example", {
    domainName: "ws-api.example.com",
    domainNameConfiguration: {
        certificateArn: aws_acm_certificate_example.arn,
        endpointType: "REGIONAL",
        securityPolicy: "TLS_1_2",
    },
});

