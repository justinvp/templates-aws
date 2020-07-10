import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directconnect.Gateway("example", {
    amazonSideAsn: "64512",
});

