---
permalink: /AC-6/
title: AC-6 - Least Privilege
---
### AWS Least Privileges  
* AWS Identity and Access Management (IAM) Policies enables organizations with many employees to create and manage multiple users under a single AWS Account. IAM policies are attached to the users, enabling centralized control of permissions for users under your AWS Account to access buckets or objects. With IAM policies, you can only grant users within your own AWS account permission to access your Amazon resources.  
* 18F AWS IAM policies are defined to grant only the required access for 18F staff necessary to perform their functions. 18F defines least privilege access to each user , group or role. and is in the planning stages to customize access  to specific resources using an authorization policy.  
* Security functions within the AWS infrastructure can explicitly be defined within IAM to include read-only permissions for any user functions.  
* 18F will incorporate running the IAM Policy Simulator to test all of its policies for least privilege access for users and groups  
  
### Cloud Foundry Least Privileges  
* Cloud Foundry uses Feature Flags which allows an administrator to turn on or off sub-sections, or features, of an application without deploying new code.  
* Currently, there are six feature flags that can be set. They are all enabled by default except user_org_creation. 1. user_org_creation: When enabled, any user can create and organization via the API. 2. private_domain_creation: When enabled, and organization manager can create private domains for that organization. 3. app_bits_upload: When enabled, space developers can upload app bits. 4. app_scaling: When enabled, space developers can perform scaling operations (i.e. change memory, disk or instances). 5. route_creation: When enabled, a space developer can create routes in a space. 6. service_instance_creation: When enabled, a space developer can create service instances in a space.  
  
