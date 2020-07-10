import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleNetworkInterfaces = pulumi.output(aws.ec2.getNetworkInterfaces({ async: true }));

export const example = exampleNetworkInterfaces.ids;

