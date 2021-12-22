По всем моим картам при попытке списать 1$ AWS получал сообщение

* Amazon is not set up to accept the CVV2 security code associated with credit cards. Your bank may be able to temporarily lift this requirement.

* The authorization is for a low dollar amount ($1.00), which your bank may decline.

* If you signed up for multiple AWS services, a $1.00 authorization may be performed for each service. Your bank may approve the first authorization and decline subsequent ones depending on their security policies.

Можете подсказать банк у которого возможно сделать платежи без CVV в интернете?

#1

>curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
>unzip awscliv2.zip
>cd aws/
>./install

>aws configure list 

          Name                    Value             Type    Location
          ----                    -----             ----    --------
       profile                <not set>             None    None
    access_key                <not set>             None    None
    secret_key                <not set>             None    None
        region                <not set>             None    None

>aws configure

>aws configure list

          Name                    Value             Type    Location
          ----                    -----             ----    --------
       profile                <not set>             None    None
    access_key     ****************TOOC shared-credentials-file
    secret_key     ****************Luuq shared-credentials-file
        region                us-west-2      config-file    ~/.aws/config

>export AWS_ACCESS_KEY_ID=(your access key id)
>export AWS_SECRET_ACCESS_KEY=(your secret access key)
>aws configure list

          Name                    Value             Type    Location
          ----                    -----             ----    --------
       profile                <not set>             None    None
    access_key     ****************TOOC              env
    secret_key     ****************Luuq              env
        region                us-west-2      config-file    ~/.aws/config


#2

###При помощи какого инструмента (из разобранных на прошлом занятии) можно создать свой образ ami?

    можно использовать CloudFormation 
