import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.pricing.getProduct({
    filters: [
        {
            field: "instanceType",
            value: "c5.xlarge",
        },
        {
            field: "operatingSystem",
            value: "Linux",
        },
        {
            field: "location",
            value: "US East (N. Virginia)",
        },
        {
            field: "preInstalledSw",
            value: "NA",
        },
        {
            field: "licenseModel",
            value: "No License required",
        },
        {
            field: "tenancy",
            value: "Shared",
        },
        {
            field: "capacitystatus",
            value: "Used",
        },
    ],
    serviceCode: "AmazonEC2",
}, { async: true }));

