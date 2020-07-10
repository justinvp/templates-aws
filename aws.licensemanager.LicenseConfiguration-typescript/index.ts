import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.licensemanager.LicenseConfiguration("example", {
    description: "Example",
    licenseCount: 10,
    licenseCountHardLimit: true,
    licenseCountingType: "Socket",
    licenseRules: ["#minimumSockets=2"],
    tags: {
        foo: "barr",
    },
});

