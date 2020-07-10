import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const by_name = pulumi.output(aws.secretsmanager.getSecret({
    name: "example",
}, { async: true }));

