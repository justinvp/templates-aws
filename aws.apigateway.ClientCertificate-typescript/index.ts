import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const demo = new aws.apigateway.ClientCertificate("demo", {
    description: "My client certificate",
});

