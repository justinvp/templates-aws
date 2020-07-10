import pulumi
import pulumi_aws as aws

kittens = aws.mediapackage.Channel("kittens",
    channel_id="kitten-channel",
    description="A channel dedicated to amusing videos of kittens.")

