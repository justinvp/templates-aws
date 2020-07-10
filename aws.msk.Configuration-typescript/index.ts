import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.msk.Configuration("example", {
    kafkaVersions: ["2.1.0"],
    serverProperties: `auto.create.topics.enable = true
delete.topic.enable = true
`,
});

