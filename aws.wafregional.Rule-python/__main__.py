import pulumi
import pulumi_aws as aws

ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[{
    "type": "IPV4",
    "value": "192.0.7.0/24",
}])
wafrule = aws.wafregional.Rule("wafrule",
    metric_name="tfWAFRule",
    predicates=[{
        "dataId": ipset.id,
        "negated": False,
        "type": "IPMatch",
    }])

