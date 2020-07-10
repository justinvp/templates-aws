import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cfg.AggregateAuthorization("example", {
    accountId: "123456789012",
    region: "eu-west-2",
});

