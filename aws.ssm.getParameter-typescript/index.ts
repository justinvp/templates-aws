import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = pulumi.output(aws.ssm.getParameter({
    name: "foo",
}, { async: true }));

