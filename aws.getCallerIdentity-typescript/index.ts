import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = pulumi.output(aws.getCallerIdentity({ async: true }));

export const accountId = current.accountId;
export const callerArn = current.arn;
export const callerUser = current.userId;

