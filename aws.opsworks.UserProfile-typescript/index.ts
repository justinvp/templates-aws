import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myProfile = new aws.opsworks.UserProfile("my_profile", {
    sshUsername: "my_user",
    userArn: aws_iam_user_user.arn,
});

