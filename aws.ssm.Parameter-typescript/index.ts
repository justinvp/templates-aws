import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ssm.Parameter("foo", {
    type: "String",
    value: "bar",
});

