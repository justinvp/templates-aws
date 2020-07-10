import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDomainIdentity = new aws.ses.DomainIdentity("example", {
    domain: "example.com",
});
const examplePolicyDocument = exampleDomainIdentity.arn.apply(arn => aws.iam.getPolicyDocument({
    statements: [{
        actions: [
            "SES:SendEmail",
            "SES:SendRawEmail",
        ],
        principals: [{
            identifiers: ["*"],
            type: "AWS",
        }],
        resources: [arn],
    }],
}, { async: true }));
const exampleIdentityPolicy = new aws.ses.IdentityPolicy("example", {
    identity: exampleDomainIdentity.arn,
    policy: examplePolicyDocument.json,
});

