# Ansible

## Table of Content:
1. [Ansible Basics](#ansible-basics)
2. [Ansible Language Basics](#ansible-language-basics)
3. [Inventory File](#inventory-file)
3. [Ansible Modules](#ansible-modules)
4. [Ansible Templates](#ansible-templates)
5. [Conditionals Registery](#conditionals-registery)

<hr>

## Ansible Basics:

- **DSL:** Domain Specific Language.
- Ansible uses YAML or JSON.
- **Agent-less**, uses SSH to connect to servers.
- Written in python, Ruby, and PowerShell.
- Sponsored by RedHat.
- **Idempotency:** Configuration management desired state.

<hr>

## Ansible Language Basics:

- Playbooks contains plays.
- Plays contain tasks.
- Tasks call modules.
- Tasks run sequentially.
- Handelers are triggered by tasks, they run once, at the end of plays.


- Playbook Example:
    ```
    ---
    - name: install and start apache
    hosts: web
    remote_user: justin
    become_method: sudo
    become_user: root
    vars:
        http_port: 80
        max_clients: 200

    tasks:
    - name: install httpd
        yum: name=httpd state=latest
    - name: write apache config file
        template: src=srv/httpd.j2 dest=/etc/httpd.conf
        notify:
        - restart: apache
    - name: start httpd
        service: name=httpd state=running

    handlers:
    - name: restart apache
        service: name=httpd state=restarted
    ```
<hr>


## Inventory File:


- Grouping hosts:
```
mail.example.com

[webservers]
foo.example.com
bar.example.com

[dbservers]
one.example.com
two.example.com
three.example.com
```


<hr>
## Ansible Templates:

- Isolate variables in template file.

```
server {
        listen   80;
        server_name  {{ SERVER_NAME }};
        root   {{ REPO_PATH }};
        location / {
                index  index.php index.html index.htm;
                autoindex on;    #enable listing of directory index
        }
}
```

- Template file:
```
- name: generate all configmap configs
  template:
      src: repo.conf.j2
      dest: /etc/nginx/conf.d/repo.conf
      owner: root
      group: root
      mode: 0744
```

<hr>

## Conditionals Registery:

- Controlling execution.

- Ansible conditionals "when" clause:

    ```
    - name: Install apache
    apt:
        name: apache2
        state: present
    when: ansible_facts['on_family'] == 'Debian'

    - name: Install apache
    yum:
        name: httpd
        state: present
    when: ansible_facts['on_family'] == 'RedHat'
    ```

