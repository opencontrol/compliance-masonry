---
permalink: /AC-3/
title: AC-3 - Access Enforcement
---
### AWS Access Enforcement  
* Security features within Amazon VPC include security groups, network ACLs, routing tables, and external gateways. Each of these items is complementary to providing a secure, isolated network that can be extended through selective enabling of direct Internet access or private connectivity to another network.  
* 18F follows AWS IAM  best practices by implementing the majority of the following-  1. Create individual accounts for anyone that requires access to the AWS infrastructure or APIs or use IAM federation from enterprise identity management system 2. Use groups or roles to assign permissions to IAM users 3. Enable multi factor authentication for all IAM users 4. Use roles for applications that run on EC2 instances 5. Delegate by using roles instead of sharing credentials 6. Rotate credentials regularly.  
* Network Access control lists (ACLs) are created to allow or deny traffic entering or exiting these subnets  
* Each subnet has routing tables attached to them to direct the flow of network traffic to Internet gateways, virtual private gateways, Network Address Translation (NAT) for private subnets or VPC peering.  
* 18F has created  specific Cloud Foundry security groups associated with VPCs to provide full control over inbound and outbound traffic.  
  
### Cloud Foundry Access Enforcement  
* Within the CF PaaS, Cloud Foundry uses security groups to act as firewalls to control outbound applications in the deployment. 18F uses Cloud Foundry ASGs to specify egress access rules for its applications. This functionality enables 18F to securely restrict application outbound traffic to predefined routes.  
* Cloud Foundry evaluates security groups and other network traffic rules in a strict priority order. Cloud Foundry returns an allow, deny, or reject result for the first rule that matches the outbound traffic request parameters, and does not evaluate any lower-priority rules.  
* Cloud Foundry implements network traffic rules using Linux iptables on the component VMs. Operators can configure rules to prevent system access from external networks and between internal components, and to restrict applications from establishing connections over the DEA network interface.  
* 18F has created a specific set of VPCs ( Live production,  staging  and Lab) for its Cloud Foundry implementation.  All VPCs have subnets used to separate and control IP address  space within each individual VPC.  Subnets must be created in order to launch Availability Zone (AZ) specific services within a VPC. 18F has setup VPC Peering between the Staging VPC and the CF Live production VPC.  
  
