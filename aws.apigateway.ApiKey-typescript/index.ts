import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDemoApiKey = new aws.apigateway.ApiKey("MyDemoApiKey", {});

