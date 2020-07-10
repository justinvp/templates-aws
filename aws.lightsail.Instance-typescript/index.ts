import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new GitLab Lightsail Instance
const gitlabTest = new aws.lightsail.Instance("gitlab_test", {
    availabilityZone: "us-east-1b",
    blueprintId: "string",
    bundleId: "string",
    keyPairName: "some_key_name",
    tags: {
        foo: "bar",
    },
});

