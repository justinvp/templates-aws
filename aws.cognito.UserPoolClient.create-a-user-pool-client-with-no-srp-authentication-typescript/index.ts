import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pool = new aws.cognito.UserPool("pool", {});
const client = new aws.cognito.UserPoolClient("client", {
    explicitAuthFlows: ["ADMIN_NO_SRP_AUTH"],
    generateSecret: true,
    userPoolId: pool.id,
});

