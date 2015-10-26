---
permalink: /NIST-800-53/AC-6/
title: AC-6 - Least Privilege
parent: NIST-800-53 Documentation
---

## Identity and Access Management
- AWS Identity and Access Management (IAM) Policies enables organizations with many employees to create and manage multiple users under a single AWS Account. IAM policies are attached to the users, enabling centralized control of permissions for users under 18F AWS Account to access services, buckets or objects. With IAM policies, 18F only grant users within its own AWS account permission to access its Amazon resources.
  - 18F AWS IAM policies are defined to grant only the required access for 18F staff necessary to perform their functions. 18F defines least privilege access to each user, group or role and is in the planning stages to customize access to specific resources using an authorization policy.
  - Security functions within the AWS infrastructure are explicitly defined within IAM to include read-only permissions for any user functions. Please see Table 9-1. AWS User Roles and Privileges for all 18F IAM accounts and roles.
  - 18F will incorporate running the IAM Policy Simulator to test all of its policies for least privilege access for users and groups.
  - AWS Roles are used to allow IAM user account and AWS services to assume only the permissions granted in order to perform required tasks.

### References

* [AWS Identity and Access Management (IAM)](https://aws.amazon.com/iam/)

### Governors

* [Roles Used by 18F](Find artifact)

* [Access Control Policy Section 3](Find artifact)

* [Account Management Flow](Find artifact)

--------

## User Account and Authentication (UAA) Server
- Cloud Foundry uses Feature Flags which allows an administrator to turn on or off sub-sections, or features, of an application without deploying new code.
- Currently, there are six feature flags that can be set. They are all enabled by default except user_org_creation.
  - `user_org_creation`&colon; When enabled, any user can create and organization via the API.
  - `private_domain_creation`&colon; When enabled, and organization manager can create private domains for that organization.
  - `app_bits_upload`&colon; When enabled, space developers can upload app bits.
  - `app_scaling`&colon; When enabled, space developers can perform scaling operations (i.e. change memory, disk or instances).
  - `route_creation`&colon; When enabled, a space developer can create routes in a space.
  - `service_instance_creation`&colon; When enabled, a space developer can create service instances in a space.
- 18F uses Orgs, Spaces, and Roles to implement least privileged access to the platform as a service. Cloud Foundry uses role-based access control (RBAC), with each role granting permissions in either an org or a space. Please see Table 9-2. Cloud Foundry User Roles and Privileges. 

### References

* [User Account and Authentication (UAA) Server](http://docs.pivotal.io/pivotalcf/concepts/architecture/uaa.html)

* [Creating and Managing Users with the UAA CLI (UAAC)](http://docs.pivotal.io/pivotalcf/adminguide/uaa-user-management.html)

* [UAA Roles](https://cf-p1-docs-prod.cfapps.io/pivotalcf/concepts/roles.html)

* [Cloud Foundry Org Access](https://github.com/cloudfoundry/cloud_controller_ng/blob/master/spec/unit/access/organization_access_spec.rb)

* [Cloud Foundry Space Access](https://github.com/cloudfoundry/cloud_controller_ng/blob/master/spec/unit/access/space_access_spec.rb)

### Governors

* [Access Control Policy Section 3](Find artifact)

* [Acccount Managment Flow](Find artifact)

--------
