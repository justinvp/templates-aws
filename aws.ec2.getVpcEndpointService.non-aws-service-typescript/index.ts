import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const custome = pulumi.output(aws.ec2.getVpcEndpointService({
    serviceName: "com.amazonaws.vpce.us-west-2.vpce-svc-0e87519c997c63cd8",
}, { async: true }));

