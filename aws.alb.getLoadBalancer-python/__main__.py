import pulumi
import pulumi_aws as aws

config = pulumi.Config()
lb_arn = config.get("lbArn")
if lb_arn is None:
    lb_arn = ""
lb_name = config.get("lbName")
if lb_name is None:
    lb_name = ""
test = aws.lb.get_load_balancer(arn=lb_arn,
    name=lb_name)

