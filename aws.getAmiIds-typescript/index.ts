import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ubuntu = pulumi.output(aws.getAmiIds({
    filters: [{
        name: "name",
        values: ["ubuntu/images/ubuntu-*-*-amd64-server-*"],
    }],
    owners: ["099720109477"],
}, { async: true }));

