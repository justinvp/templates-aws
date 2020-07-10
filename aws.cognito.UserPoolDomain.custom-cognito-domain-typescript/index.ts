import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cognito.UserPool("example", {});
const main = new aws.cognito.UserPoolDomain("main", {
    certificateArn: aws_acm_certificate_cert.arn,
    domain: "example-domain.example.com",
    userPoolId: example.id,
});

