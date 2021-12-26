#1

#2

>terraform workspace list

    * default

>terraform workspace new stage

    Created and switched to workspace "stage"!
    
    You're now on a new, empty workspace. Workspaces isolate their state,
    so if you run "terraform plan" Terraform will not see any existing state
    for this configuration.

> terraform workspace new prod

    Created and switched to workspace "prod"!
    
    You're now on a new, empty workspace. Workspaces isolate their state,
    so if you run "terraform plan" Terraform will not see any existing state
    for this configuration.

>terraform workspace list

      default
    * prod
      stage

Переключение на STAGE

> terraform workspace select stage
    
    Switched to workspace "stage".

>terraform workspace list

      default
      prod
    * stage

Переключение на PROD

> terraform workspace select prod
    
    Switched to workspace "prod".

>terraform workspace list

      default
    * prod
      stage


>terraform plan
        
        An execution plan has been generated and is shown below.
        Resource actions are indicated with the following symbols:
          + create
        
        Terraform will perform the following actions:
        
          # aws_s3_bucket.bucket will be created
          + resource "aws_s3_bucket" "bucket" {
              + acceleration_status         = (known after apply)
              + acl                         = "private"
              + arn                         = (known after apply)
              + bucket                      = "netology-bucket-prod"
              + bucket_domain_name          = (known after apply)
              + bucket_regional_domain_name = (known after apply)
              + force_destroy               = false
              + hosted_zone_id              = (known after apply)
              + id                          = (known after apply)
              + region                      = (known after apply)
              + request_payer               = (known after apply)
              + tags                        = {
                  + "Environment" = "prod"
                  + "Name"        = "Bucket1"
                }
              + website_domain              = (known after apply)
              + website_endpoint            = (known after apply)
        
              + versioning {
                  + enabled    = (known after apply)
                  + mfa_delete = (known after apply)
                }
            }
        
        Plan: 1 to add, 0 to change, 0 to destroy.