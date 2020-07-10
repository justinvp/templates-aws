import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const filter = new aws.ses.ReceiptFilter("filter", {
    cidr: "10.10.10.10",
    policy: "Block",
});

