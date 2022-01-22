#1
    
resource = https://github.com/hashicorp/terraform-provider-aws/blob/d87b1fca5d57a6fc49c0647657c11eff14e4ab6c/internal/provider/provider.go#L741

data_source = https://github.com/hashicorp/terraform-provider-aws/blob/d87b1fca5d57a6fc49c0647657c11eff14e4ab6c/internal/provider/provider.go#L345

#2

###1
    
    "name": {
                Type:          schema.TypeString,
                Optional:      true,
                Computed:      true,
                ForceNew:      true,
                ConflictsWith: []string{"name_prefix"},
            }

Строка кода https://github.com/hashicorp/terraform-provider-aws/blob/8c31926c699e8db4f5b1bc258a44dc376c38aedc/internal/service/sqs/queue.go#L82 

###2

Название машины https://github.com/hashicorp/terraform-provider-aws/blob/a7cd46f0c7e33b5e996e8b5b86ec764682e4bc11/internal/service/sfn/state_machine.go#L72

    "name": {
                    Type:         schema.TypeString,
                    Required:     true,
                    ForceNew:     true,
                    ValidateFunc: validStateMachineName,
                }

Функция валидация имени ValidateFunc: validStateMachineName
https://github.com/hashicorp/terraform-provider-aws/blob/a7cd46f0c7e33b5e996e8b5b86ec764682e4bc11/internal/service/sfn/validate.go#L8

    func validStateMachineName(v interface{}, k string) (ws []string, errors []error) {
        value := v.(string)
        if len(value) > 80 {
            errors = append(errors, fmt.Errorf("%q cannot be longer than 80 characters", k))
        }
    
        if !regexp.MustCompile(`^[a-zA-Z0-9-_]+$`).MatchString(value) {
            errors = append(errors, fmt.Errorf(
                "%q must be composed with only these characters [a-zA-Z0-9-_]: %v", k, value))
        }
        return
    }

Длина имени не может быть более 80 символов


###3

Регулярное выражение для имени ^[a-zA-Z0-9-_]+$
Имя должно начинаться только с букв латиницы (верхнего и нижнего регистра), цифры, подчеркивания


#2

Скомпилировал самый простой вариант провайдера. Далее буду разбираться как улучшать провайдер.

    oracle@coffe:~/terraform-provider-hashicups$ go build -o terraform-provider-hashicups
    oracle@coffe:~/terraform-provider-hashicups$ ./terraform-provider-hashicups
    This binary is a plugin. These are not meant to be executed directly.
    Please execute the program that consumes these plugins, which will
    load any plugins automatically

Сам провайдер можно найти [тут](terraform-provider-hashicups_v1)

