import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConnection = new aws.directconnect.Connection("example", {
    bandwidth: "1Gbps",
    location: "EqSe2",
});
const exampleLinkAggregationGroup = new aws.directconnect.LinkAggregationGroup("example", {
    connectionsBandwidth: "1Gbps",
    location: "EqSe2",
});
const exampleConnectionAssociation = new aws.directconnect.ConnectionAssociation("example", {
    connectionId: exampleConnection.id,
    lagId: exampleLinkAggregationGroup.id,
});

