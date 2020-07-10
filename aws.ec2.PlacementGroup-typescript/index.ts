import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const web = new aws.ec2.PlacementGroup("web", {
    strategy: "cluster",
});

