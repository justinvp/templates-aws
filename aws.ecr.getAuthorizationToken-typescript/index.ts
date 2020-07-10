import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const token = pulumi.output(aws.ecr.getAuthorizationToken({ async: true }));

