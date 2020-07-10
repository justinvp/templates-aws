import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appsync.GraphQLApi("example", {
    authenticationType: "OPENID_CONNECT",
    openidConnectConfig: {
        issuer: "https://example.com",
    },
});

