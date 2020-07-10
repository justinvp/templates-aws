import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.User("example", {
    email: "author@example.com",
    identityType: "IAM",
    userName: "an-author",
    userRole: "AUTHOR",
});

