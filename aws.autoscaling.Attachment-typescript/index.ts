import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new load balancer attachment
const asgAttachmentBar = new aws.autoscaling.Attachment("asg_attachment_bar", {
    autoscalingGroupName: aws_autoscaling_group_asg.id,
    elb: aws_elb_bar.id,
});

