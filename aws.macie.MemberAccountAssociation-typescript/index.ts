import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.macie.MemberAccountAssociation("example", {
    memberAccountId: "123456789012",
});

