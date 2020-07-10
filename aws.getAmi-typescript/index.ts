import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.getAmi({
    executableUsers: ["self"],
    filters: [
        {
            name: "name",
            values: ["myami-*"],
        },
        {
            name: "root-device-type",
            values: ["ebs"],
        },
        {
            name: "virtualization-type",
            values: ["hvm"],
        },
    ],
    mostRecent: true,
    nameRegex: "^myami-\\d{3}",
    owners: ["self"],
}, { async: true }));

