import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("example", {});
const current = pulumi.output(aws.getRegion({ async: true }));
const exampleProductSubscription = new aws.securityhub.ProductSubscription("example", {
    productArn: pulumi.interpolate`arn:aws:securityhub:${current.name!}:733251395267:product/alertlogic/althreatmanagement`,
}, { dependsOn: [exampleAccount] });

