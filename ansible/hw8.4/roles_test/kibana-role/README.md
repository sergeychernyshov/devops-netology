Kibana
=========

Simple download binaries from official website and install Kibana.

Role Variables
--------------

There is only two variables that you can redefine in your playbook.

```yaml
kibana_version: "7.14.0" # Use for download only this version of kibana
kibana_home: "/opt/kibana/{{ kibana_version }}" # Use for unpackage distro and create KIBANA_HOME variable
```

Example Playbook
----------------

Including an example of how to use your role (for instance, with variables passed in as parameters) is always nice for users too:

```yaml
- hosts: kibana
  roles:
      - kibana
```

License
-------

BSD

Author Information
------------------

Sergey Chernyshov