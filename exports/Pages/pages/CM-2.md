---
permalink: /NIST-800-53/CM-2/
title: CM-2 - Baseline Configuration
parent: NIST-800-53 Documentation
---

## AlienVault
AlienVault USM for AWS is provided by the vendor as a secure hardened AMI image that is deployed using a cloudformation template.
### References

* [AlienVault Website](http://www.alienvault.com)

--------

## Amazon Machine Images
- Linux instances are based on the standard AWS AMI images with configuration to GSA requirements based on secure configurations documented in CM-6.
- AlienVault USM for AWS is provided by the vendor as a secure hardened AMI image that is deployed using a cloudformation template.
- NIST guidance, best practices, CIS benchmarks along with standard and hardened Operating System AMIs have been utilized.
- DevOps maintain copies of the latest Production Software Baseline, which includes the following elements: Manufacturer, Type, Version number, Software, Databases, and Stats.

### References

* [Amazon Machine Images](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIs.html)

--------

## AWS Service Catalog
AWS Service Catalog allows 18F to centrally manage commonly deployed IT services, and helps achieve consistent governance and meet compliance requirements, while enabling users to quickly deploy only the approved IT services they need.
### References

* [AWS Service Catalog Documentation](https://aws.amazon.com/servicecatalog/)

--------

## Cloud Formation
- DevOps maintain baseline configurations for VPC, EBS, EC2 instances and AMIs. AWS Cloud Formation templates help 18F maintain a strict configuration management scheme of the cloud infrastructure. If an error or misconfiguration of the infrastructure or associated security mechanism (security groups, NACLs) is detected, the administrators can analyze the current infrastructure templates; compare with previous versions, and redeploy the configurations to a known and approved state.
- AWS Cloud Formation templates are the approved baseline for all changes to the infrastructure and simplify provisioning and management on AWS. They provide an automated method to assess the status of an operational infrastructure against an approved baseline.

### References

* [What is AWS CloudFormation?](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html)

--------

## Amazon Elastic Block Store
DevOps maintain baseline configurations for VPC, EBS, EC2 instances and AMIs. AWS Cloud Formation templates help 18F maintain a strict configuration management scheme of the cloud infrastructure. If an error or misconfiguration of the infrastructure or associated security mechanism (security groups, NACLs) is detected, the administrators can analyze the current infrastructure templates; compare with previous versions, and redeploy the configurations to a known and approved state.
### References

* [Amazon Elastic Block Store](https://aws.amazon.com/ebs/)

--------

## Amazon Elastic Compute Cloud
DevOps maintain baseline configurations for VPC, EBS, EC2 instances and AMIs. AWS Cloud Formation templates help 18F maintain a strict configuration management scheme of the cloud infrastructure. If an error or misconfiguration of the infrastructure or associated security mechanism (security groups, NACLs) is detected, the administrators can analyze the current infrastructure templates; compare with previous versions, and redeploy the configurations to a known and approved state.
### References

* [Amazon Elastic Compute Cloud](https://aws.amazon.com/ec2/)

--------

## Manifests
Configure UAA clients and users using a standard BOSH manifest for cloud Foundry Deployment, Limit and manage these clients and users as you would any other kind of privileged account.
### References

* [Deploying with Application Manifests](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html)

--------
