import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleResourceShare = new aws.ram.ResourceShare("example", {
    // ... other configuration ...
    allowExternalPrincipals: true,
});
const examplePrincipalAssociation = new aws.ram.PrincipalAssociation("example", {
    principal: "111111111111",
    resourceShareArn: exampleResourceShare.arn,
});

