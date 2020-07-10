import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testStaticIp = new aws.lightsail.StaticIp("test", {});
const testInstance = new aws.lightsail.Instance("test", {
    availabilityZone: "us-east-1b",
    blueprintId: "string",
    bundleId: "string",
    keyPairName: "some_key_name",
});
const testStaticIpAttachment = new aws.lightsail.StaticIpAttachment("test", {
    instanceName: testInstance.id,
    staticIpName: testStaticIp.id,
});

