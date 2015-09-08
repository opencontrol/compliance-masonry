---
permalink: /CM-6/
title: CM-6 - Configuration Settings
---
### AWS documented configuration settings  
* System Administrators / Security Engineers maintain the baseline configuration for VPC, EBS.  Best practices are utilized as there are no benchmarks available.  
* 18F incorporates the AWS Service Catalog to create, manage, and distribute catalogs of approved products to end users, who can then access the products they need in a personalized portal. System administrators can control which users have access to each application or AWS resource to enforce compliance with organizational business policies for documenting, implementing, monitoring and controlling changes to configuration on their information systems built on the AWS infrastructure.  
* GSA and 18F follow industry best practices and guidance provided in NIST Special Publication 800-70, Security Configuration Checklist Program for IT Products.  
* The organization uses FISMA compliant and hardened AMIs within its AWS infrastructure.  
  
### Cloud Foundry documented configuration settings  
* 18F follows Cloud Foundry best practices for configuring and implementing the platform as a service  
* Cloud Foundry configuration settings are documented within the deployment manifest on the 18F Github and Cloud Foundry sites.  
* The following are approved baseline configuration settings related to the Cloud Foundry platform as a service. * The 18F CloudCoundry manifest can be found at https://github.com/18F/cloud-foundry-manifests/blob/master/cf/cf-jobs.yml#L134-L557 * 18F BOSH Stemcells are used for the standard baseline OS images and software vulnerabilty management updates found at https://github.com/cloudfoundry/bosh/blob/master/bosh-stemcell/OS_IMAGES.md * Updates to new BOSH stemcells are located and stored within Amazon S3 http://boshartifacts.cloudfoundry.org/file_collections?type=stemcells *The specifications of the current release of BOSH stemcells are located on Github https://github.com/cloudfoundry/bosh/blob/master/bosh-stemcell/spec/stemcells/ubuntu_trusty_spec.rb *Other documentation related to Cloud Foundry can be found within the 18F Github Repository located at https://docs.18f.gov/ops/repos/.  
  
### Cloud Foundry Configuration Settings implementations  
* 18F DevOps use AWS pre-configured baseline cloudformation templates hardened to meet and Cloud Foundry best practices to deploy the platform as a service  
* DevOps implements Cloud Foundry standard BOSH stemcells for baseline OS configuration located at  https://github.com/cloudfoundry/bosh/blob/master/bosh-stemcell/OS_IMAGES.md  
* DevOps implements manifest templates written in yml to automate deployment of multiple applications at once and the platform within AWS with consistency and reproducibility. All 18F cloud foundry manifest templates are located on the 18F Github site  https://github.com/18F/cloud-foundry-manifests  
  
