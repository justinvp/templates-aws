import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = pulumi.output(aws.getAvailabilityZones({ async: true }));
const currentRegion = pulumi.output(aws.getRegion({ async: true }));
const currentCallerIdentity = pulumi.output(aws.getCallerIdentity({ async: true }));
const fooEip = new aws.ec2.Eip("foo", {
    vpc: true,
});
const fooProtection = new aws.shield.Protection("foo", {
    resourceArn: pulumi.interpolate`arn:aws:ec2:${currentRegion.name!}:${currentCallerIdentity.accountId}:eip-allocation/${fooEip.id}`,
});

