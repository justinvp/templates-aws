import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const userUpdates = new aws.sns.Topic("user_updates", {});

