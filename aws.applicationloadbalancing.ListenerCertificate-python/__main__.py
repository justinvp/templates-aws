import pulumi
import pulumi_aws as aws

example_certificate = aws.acm.Certificate("exampleCertificate")
front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
front_end_listener = aws.lb.Listener("frontEndListener")
example_listener_certificate = aws.lb.ListenerCertificate("exampleListenerCertificate",
    certificate_arn=example_certificate.arn,
    listener_arn=front_end_listener.arn)

