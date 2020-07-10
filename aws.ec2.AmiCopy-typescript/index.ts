import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.AmiCopy("example", {
    description: "A copy of ami-xxxxxxxx",
    sourceAmiId: "ami-xxxxxxxx",
    sourceAmiRegion: "us-west-1",
    tags: {
        Name: "HelloWorld",
    },
});

