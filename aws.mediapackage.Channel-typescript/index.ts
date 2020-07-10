import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const kittens = new aws.mediapackage.Channel("kittens", {
    channelId: "kitten-channel",
    description: "A channel dedicated to amusing videos of kittens.",
});

