import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const user1 = new aws.iam.User("user1", {});
const group1 = new aws.iam.Group("group1", {});
const group2 = new aws.iam.Group("group2", {});
const example1 = new aws.iam.UserGroupMembership("example1", {
    groups: [
        group1.name,
        group2.name,
    ],
    user: user1.name,
});
const group3 = new aws.iam.Group("group3", {});
const example2 = new aws.iam.UserGroupMembership("example2", {
    groups: [group3.name],
    user: user1.name,
});

