import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleUser = new aws.iam.User("example", {
    forceDestroy: true,
    path: "/",
});
const exampleUserLoginProfile = new aws.iam.UserLoginProfile("example", {
    pgpKey: "keybase:some_person_that_exists",
    user: exampleUser.name,
});

export const password = exampleUserLoginProfile.encryptedPassword;

