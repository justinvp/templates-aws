import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const centos = pulumi.output(aws.ssm.getPatchBaseline({
    namePrefix: "AWS-",
    operatingSystem: "CENTOS",
    owner: "AWS",
}, { async: true }));

