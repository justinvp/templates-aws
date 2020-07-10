import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleNetworkAcls = pulumi.output(aws.ec2.getNetworkAcls({
    vpcId: var_vpc_id,
}, { async: true }));

export const example = exampleNetworkAcls.ids;

