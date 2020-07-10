import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultOpenIdConnectProvider = new aws.iam.OpenIdConnectProvider("default", {
    clientIdLists: ["266362248691-342342xasdasdasda-apps.googleusercontent.com"],
    thumbprintLists: [],
    url: "https://accounts.google.com",
});

