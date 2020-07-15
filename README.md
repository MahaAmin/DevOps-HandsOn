# DevOps Task 

## Table Of Content:
1. [Studying Docker](Docker.md)
2. [Dockerfile](#dockerfile)
3. [Ansible](#ansible)
4. [Monitoring Tools](#monitoring-tools)
5. [Locust Load Test](#locust-load-test)

<hr>

## Dockerfile

Multistaged Dockerfile :
1. Stage #1: golang:1.13-alpine image to build the be-svc and fe-svc packages, and produce 2 binary files: be-bin and fe-bin.

2. Stage #2: alpine:3.11 image to run be-bin.

<hr>

## Ansible

[Ansible Study Book](Ansible.md)

<hr>

## Monitoring Tools:

[Monitoring Tools Study Book](Monitoring.md)

<hr>

## Locust Load Test:

[Getting Started With Locust Load Testing](locust/Locust-Load-Test.md)