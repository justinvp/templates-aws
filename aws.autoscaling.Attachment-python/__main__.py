import pulumi
import pulumi_aws as aws

# Create a new load balancer attachment
asg_attachment_bar = aws.autoscaling.Attachment("asgAttachmentBar",
    autoscaling_group_name=aws_autoscaling_group["asg"]["id"],
    elb=aws_elb["bar"]["id"])

