import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appsync.GraphQLApi("example", {
    additionalAuthenticationProviders: [{
        authenticationType: "AWS_IAM",
    }],
    authenticationType: "API_KEY",
});

