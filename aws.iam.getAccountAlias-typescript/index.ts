import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = pulumi.output(aws.iam.getAccountAlias({ async: true }));

export const accountId = current.accountAlias;

