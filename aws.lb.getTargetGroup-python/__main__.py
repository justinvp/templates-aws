import pulumi
import pulumi_aws as aws

config = pulumi.Config()
lb_tg_arn = config.get("lbTgArn")
if lb_tg_arn is None:
    lb_tg_arn = ""
lb_tg_name = config.get("lbTgName")
if lb_tg_name is None:
    lb_tg_name = ""
test = aws.lb.get_target_group(arn=lb_tg_arn,
    name=lb_tg_name)

