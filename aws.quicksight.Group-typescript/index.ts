import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.Group("example", {
    groupName: "tf-example",
});

