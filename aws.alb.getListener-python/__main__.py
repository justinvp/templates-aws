import pulumi
import pulumi_aws as aws

config = pulumi.Config()
listener_arn = config.require_object("listenerArn")
listener = aws.lb.get_listener(arn=listener_arn)
selected = aws.lb.get_load_balancer(name="default-public")
selected443 = aws.lb.get_listener(load_balancer_arn=selected.arn,
    port=443)

