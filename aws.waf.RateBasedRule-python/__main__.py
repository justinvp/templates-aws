import pulumi
import pulumi_aws as aws

ipset = aws.waf.IpSet("ipset", ip_set_descriptors=[{
    "type": "IPV4",
    "value": "192.0.7.0/24",
}])
wafrule = aws.waf.RateBasedRule("wafrule",
    metric_name="tfWAFRule",
    predicates=[{
        "dataId": ipset.id,
        "negated": False,
        "type": "IPMatch",
    }],
    rate_key="IP",
    rate_limit=100,
    opts=ResourceOptions(depends_on=["aws_waf_ipset.ipset"]))

