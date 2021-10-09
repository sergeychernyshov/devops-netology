#1

    Не возможно вычислить значение "c" так как переменные a и b разного типа. 
    получить 12 можно следующим образом
    a = 1
    b = '2'
    c = str(a) +b

    получить 3 можно следующим образом
    c = a + int(b)

#2
    
    import os

    pathRepo = "~/netology/sysadm-homeworks"
    result_full_path = os.popen("cd " + pathRepo + " && " + "git rev-parse --show-toplevel").read()
    prepare_result_full_path = result_full_path.replace('\n', '')
    result_os = os.popen("cd " + pathRepo + " && " + "git status").read()
    for result in result_os.split('\n'):
        if result.find('modified') != -1:
            prepare_result = result.replace('\tmodified:   ', '')
            print(prepare_result_full_path + "/" + prepare_result)

#3

    import os
    import sys
    
    pathRepo = sys.argv[1]
    if os.path.isdir(pathRepo+"\\.git"):
        result_full_path = os.popen("cd " + pathRepo + " && " + "git rev-parse --show-toplevel").read()
        prepare_result_full_path = result_full_path.replace('\n', '')
        result_os = os.popen("cd " + pathRepo + " && " + "git status").read()
        for result in result_os.split('\n'):
            if result.find('modified') != -1:
                prepare_result = result.replace('\tmodified:   ', '')
                print(prepare_result_full_path + "/" + prepare_result)
    else:
        print("directory \"" + pathRepo + "\" not exist git repo")

#4

    import socket
    import os.path
    import json
    
    
    def write_data_to_file(file_name, json_data):
        with open(file_name, 'w') as outfile:
            json.dump(json_data, outfile)
    
    
    def read_data_from_file(file_name):
        with open(file_name) as json_data:
            data_file = json.load(json_data)
        return data_file
    
    
    def check_ip(sd, sd_last, services_list):
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
    services = ["drive.google.com", "mail.google.com", "google.com"]
    services_data = {}
    for service in services:
        service_ip_address = socket.gethostbyname(service)
        services_data[service] = service_ip_address
        print(service + " " + service_ip_address)
    
    if os.path.isfile(services_data_file):
        services_data_last_check = read_data_from_file(services_data_file)
        result = check_ip(services_data, services_data_last_check, services)
        if not result:
            write_data_to_file(services_data_file, services_data)
    else:
        write_data_to_file(services_data_file, services_data)
    
    


