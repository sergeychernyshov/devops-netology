#1

    * "ip : 71.78.22.43 - пропущены кавычки после ip
    * "ip : 71.78.22.43 - 71.78.22.43 строка без кавычек    

    Правильный скрипт:
    { "info" : "Sample JSON output from our service\t",
        "elements" :[
            { "name" : "first",
            "type" : "server",
            "ip" : 7175 
            },
            { "name" : "second",
            "type" : "proxy",
            "ip" : "71.78.22.43"
            }
        ]
    }

#2

    import socket
    import os.path
    import json
    import yaml
    
    
    def write_data_to_file(file_name, yaml_file, json_data):
        with open(file_name, 'w') as outfile:
            json.dump(json_data, outfile)
        with open(yaml_file, 'w') as file:
            yaml.dump(json_data, file)
    
    
    def read_data_from_file(file_name):
        with open(file_name) as json_data:
            data_file = json.load(json_data)
        return data_file
    
    
    def check_ip(sd, sd_last):
        check_result = True
        union_list = list(sd | sd_last)
        for service_check in union_list:
            if service_check in sd and service_check in sd_last:
                if sd[service_check] != sd_last[service_check]:
                    print('')
                    print("ERROR: " + service_check + " IP mismatch: " + sd_last[service_check] + " " + sd[service_check])
                    check_result = False
            else:
                if service_check in sd:
                    print('')
                    print("NEW SERVICE: " + service_check + " " + sd[service_check])
                else:
                    print('')
                    print("DELETE SERVICE: " + service_check + " " + sd_last[service_check])
                check_result = False
        return check_result
    
    
    services_data_file = "services_data.txt"
    services_yaml_file = "services_yaml.txt"
    services = ["drive.google.com", "mail.google.com", "google.com"]
    services_data = {}
    for service in services:
        service_ip_address = socket.gethostbyname(service)
        services_data[service] = service_ip_address
        print(service + " " + service_ip_address)
    
    if os.path.isfile(services_data_file):
        services_data_last_check = read_data_from_file(services_data_file)
        result = check_ip(services_data, services_data_last_check)
        if not result:
            write_data_to_file(services_data_file, services_yaml_file, services_data)
    else:
        write_data_to_file(services_data_file, services_yaml_file, services_data)
