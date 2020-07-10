import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.transfer.getServer({
    serverId: "s-1234567",
}, { async: true }));

