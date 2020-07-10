import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const by_arn = pulumi.output(aws.secretsmanager.getSecret({
    arn: "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456",
}, { async: true }));

