import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const portfolio = new aws.servicecatalog.Portfolio("portfolio", {
    description: "List of my organizations apps",
    providerName: "Brett",
});

