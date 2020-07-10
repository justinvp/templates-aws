import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloud9.EnvironmentEC2("example", {
    instanceType: "t2.micro",
});

