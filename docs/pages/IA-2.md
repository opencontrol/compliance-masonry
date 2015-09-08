---
permalink: /IA-2/
title: IA-2 - User Identification and Authentication
---
### AWS User Identification and Authentication  
* All users have individually unique identifiers to access and authenticates to the environment.  
* 18F AWS IAM users are placed into groups based on their assigned roles and permissions.  
* Additional temporary permission are delegated with the IAM roles usually for applications that run on EC2 Instances in order to access AWS resources (i.e. Amazon S3 buckets, DynamoDB data)  
* All user accounts for 18F staff are maintained within the 18F AWS Environment.  
* Shared or group authenticators are not utilized, Service accounts are implemented as Managed Services Accounts within AWS.  
  
### Cloud Foundry User Identification and Authentication  
* The UAA is the identity management service for Cloud Foundry. Its primary role is as an OAuth2 provider, issuing tokens for client applications to use when they act on behalf of Cloud Foundry users. In collaboration with the login server, it authenticate users with their Cloud Foundry credentials, and can act as an SSO service using those credentials (or others). It has endpoints for managing user accounts and for registering OAuth2 clients, as well as various other management functions.  
* All users have individually unique identifiers to access and authenticates to the environment  
* Shared or group authenticators are not utilized, with the exception of root administrative and emergency administrative accounts.  
  
