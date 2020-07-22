# CKA: Certified Kuberntes Administrator

## Table of Content:

1. [Core Concepts](#core-concepts)
    1. [Cluster Architecture](#cluster-architecture)
    2. [API Primitives](#api-primitives)
    3. [Services and Other Network Primitives](#services-and-other-network-primitives)
2. [Scheduling](#scheduling)
    1. [Labels And Selectors](#labels-and-selectors)
    2. [Resource Limits](#resource-limits)
    3. [Manual Scheduling](#manual-scheduling)
    4. [Daemon Sets](#daemon-sets)
    5. [Multiple Schedulers](#multiple-schedulers)
    6. [Scheduler Events](#scheduler-events)
    7. [Configure Kubernetes Scheduler](#configure-kubernetes-scheduler)
3. [Logging Monitoring](#logging-monitoring)
    1. [Monitor Cluster Components](#monitor-cluster-components)
    2. [Monitor Applications](#monitor-applications)
    3. [Monitor Cluster Component Logs](#monitor-cluster-component-logs)
    4. [Application Logs](#application-logs)
4. [Application Life Cycle](#application-life-cycle)
    1. [Rolling Updates And Rollbacks In Deploy](#rolling-updates-and-rollbacks-in-deploy)
    2. [Configure Applications](#configure-applications)
    3. [Scale Applications](#scale-applications)
    4. [Self-Healing Applications](#self-healing-applications)
5. [Cluster Maintenance](#cluster-maintenance)
    1. [Cluster Upgrade Process](#cluster-upgrade-process)
    2. [Operating System Upgrades](#operating-system-upgrades)
    3. [Backup And Restore Methodologies](#backup-and-restore-methodologies)
6. [Security](#security)
    1. [Authentication And Authorization](authentication-and-authorization)
    2. [Kubernetes Security](#kubernetes-security)
    3. [Images Securely](#images-securely)
    4. [Network Polices](#network-polices)
    5. [TLS Certificates For Cluster Components](#tls-certificates-for-cluster-components)
    6. [Security Contexts](#security-contexts)
    7. [Secure Persistance Value Store](#secure-persistance-value-store)
7. [Storage](#storage)
    1. [Persistent Volumes](#persistent-volumes)
    2. [Persistent Volume Claims](#persistent-volume-claims)
    3. [Configure Applications With Persistent Storage](#configure-applications-with-persistent-storage)
    4. [Access Models For Volumes](#access-models-for-volumes)
    5. [Kubernetes Storage Object](#kubernetes-storage-object)
8. [Networking](#networking)
    1. [Prerequisites](#prerequisites)
    2. [Networking Configuration on Cluster Nodes](#networking-configuration-on-cluster-nodes)
    3. [Service Networking](#service-networking)
    4. [POD Networking Concepts](#pod-networking-concepts)
    5. [Network Loadbalancer](#network-loadbalancer)
    6. [Ingress](#ingress)
    7. [Cluster DNS](#cluster-dns)
    8. [CNI](#cni)
9. [Installation Configuration Validation](#installation-configuration-validation)
    1. [Design a Kubernetes Cluster](#design-a-kubernetes-cluster)
    2. [Install Kubernetes Master and Nodes](#install-kubernetes-master-and-nodes)
    3. [Secure Cluster Communication](#secure-cluster-communication)
    4. [HA Kubernetes Cluster](#ha-kubernetes-cluster)
    5. [Provision Infrastructure](#provision-infrastructure)
    6. [Choose A Network Solution](#choose-a-network-solution)
    7. [Run And Analyze End-To-End Test](run-and-analyze-end-to-end-test)
    8. [Node End-To-End Tests](#node-end-to-end-tests)
10. [Troubleshooting](#troubleshooting)
    1. [Application Failure](#application-failure)
    2. [Worker Node Failure](#worker-node-failure)
    3. [Control Plane Failure](#control-plane-failure)
    4. [Networking](#networking)

<hr>


## Core Concepts

### Cluster Architecture

- Master Node:
    - ETCD Cluster.
    - Kupe API Server.
    - Kube Controller Manager.
    - Kube Scheduler.
    - Cloud Controller Manager.

- Worker Nodes:
    - Kubelet
    - Kube Proxy
    - Container Runtime Environment.

![Kubernetes Architecture](img/kubernetes-architecture.png)

<hr>

### API Primitives

**ETCD:**

- Distributed key-value store.
- ETCD listens to port 2379 by default.
- To set key-value pair:
    ```
    $ ./etcdctl set <key> <value>  
    ```

- To get tha value stored with a key:
    ```
    $ ./etcdctl get <key>
    ```

- ETCD in kubernetes:
    - Stores information about the cluster as: nodes, pods, configs, secrets, accounts, roles, bindings, ...etc.

    - All information from **kubectl get** are retrieved from etcd server.

    - Every change made on the cluster gets updated in the etcd.
