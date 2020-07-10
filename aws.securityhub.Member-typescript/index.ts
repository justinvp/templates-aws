import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("example", {});
const exampleMember = new aws.securityhub.Member("example", {
    accountId: "123456789012",
    email: "example@example.com",
    invite: true,
}, { dependsOn: [exampleAccount] });

