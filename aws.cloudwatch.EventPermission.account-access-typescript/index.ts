import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const devAccountAccess = new aws.cloudwatch.EventPermission("DevAccountAccess", {
    principal: "123456789012",
    statementId: "DevAccountAccess",
});

