import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.iam.getPolicy({
    arn: "arn:aws:iam::123456789012:policy/UsersManageOwnCredentials",
}, { async: true }));

