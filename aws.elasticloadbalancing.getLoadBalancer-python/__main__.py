import pulumi
import pulumi_aws as aws

config = pulumi.Config()
lb_name = config.get("lbName")
if lb_name is None:
    lb_name = ""
test = aws.elb.get_load_balancer(name=lb_name)

