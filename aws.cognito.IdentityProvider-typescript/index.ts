import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cognito.UserPool("example", {
    autoVerifiedAttributes: ["email"],
});
const exampleProvider = new aws.cognito.IdentityProvider("example_provider", {
    attributeMapping: {
        email: "email",
        username: "sub",
    },
    providerDetails: {
        authorize_scopes: "email",
        client_id: "your client_id",
        client_secret: "your client_secret",
    },
    providerName: "Google",
    providerType: "Google",
    userPoolId: example.id,
});

