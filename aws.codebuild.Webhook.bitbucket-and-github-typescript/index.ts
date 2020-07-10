import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codebuild.Webhook("example", {
    filterGroups: [{
        filters: [
            {
                pattern: "PUSH",
                type: "EVENT",
            },
            {
                pattern: "master",
                type: "HEAD_REF",
            },
        ],
    }],
    projectName: aws_codebuild_project_example.name,
});

