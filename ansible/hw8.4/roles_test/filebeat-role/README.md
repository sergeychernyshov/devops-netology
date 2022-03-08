Filebeat
=========

Simple download binaries from official website and install Filebeat.

Role Variables
--------------

There is only two variables that you can redefine in your playbook.

```yaml
filebeat_version: "7.14.0" # Use for download only this version of filebead
filebeat_home: "/opt/kibana/{{ elastic_version }}" # Use for unpackage distro and create FILEBEAT_HOME variable
```

Example Playbook
----------------

Including an example of how to use your role (for instance, with variables passed in as parameters) is always nice for users too:

```yaml
- hosts: app
  roles:
      - filebeat
```

License
-------

BSD

Author Information
------------------

Sergey Chernyshov