import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const group = new aws.iam.Group("group", {});
const userOne = new aws.iam.User("user_one", {});
const userTwo = new aws.iam.User("user_two", {});
const team = new aws.iam.GroupMembership("team", {
    group: group.name,
    users: [
        userOne.name,
        userTwo.name,
    ],
});

