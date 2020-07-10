import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentUser = pulumi.output(aws.getCanonicalUserId({ async: true }));
const bucket = new aws.s3.Bucket("bucket", {
    grants: [
        {
            id: currentUser.id,
            permissions: ["FULL_CONTROL"],
            type: "CanonicalUser",
        },
        {
            permissions: [
                "READ",
                "WRITE",
            ],
            type: "Group",
            uri: "http://acs.amazonaws.com/groups/s3/LogDelivery",
        },
    ],
});

