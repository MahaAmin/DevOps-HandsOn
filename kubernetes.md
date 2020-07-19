# Kubernetes

## Table of Content:

1. [Kubernetes For The Absolute Beginners](#kubernetes-for-the-absolute-beginners)
    1. [Kubernetes Overview](#kubernetes-overview)
    2. [Docker Containers](#docker-containers)
    3. [Containers Orchestration](#containers-orchestration)
    4. [Kubernetes Architecture](#kubernetes-architecture)
    4. [Kubernetes Setup](#kubernetes-setup)
    5. [Kubernetes Concepts (PODs - ReplicaSets - Deployment - Services)](#kubernetes-concepts)
    6. [Networking in Kubernetes](#networking-in-kubernetes)
    7. [KubeCtl](#kubectl)
    8. [YAML](#yaml)
    9. [Kubernetes on Cloud](#kubernetes-on-cloud)
2. [Kubernetes Certificates](#kubernetes-certificates)
3. [Resources](#resources)


<hr>


## Kubernetes For The Absolute Beginners:

### Kubernetes Overview:

- Built by Google.
- It is now open source project.

<hr>


### Docker Containers:

- Refer to [Docker Study Book](Docker.md).

<hr>

### Containers Orchestration:

- Kubernetes is an orchestration technlogy.
- There are several other technologies for orchestration like: **Docker Swarm** and **MESOS** from apache.
- Kubernetes is now supported by all cloud providers.
- One of the most ranked os projects on GitHub.

<hr>


### Kubernetes Architecture:

![kubernetes-architecture](img/kubernetes-architecture.png)

- **Nodes (Minios):** is a machine physical or virtual.
- **Cluster:** set of nodes grouped together.
- **Master:** a node where kubernetes is installed, and configured as master. Responsible of the actual orchestration of containers in the worker nodes.
- **Container Runtime**: Docker, 

<hr>

## Kubernetes Certificates:

- **CKA** : Certified Kubernetes Adminstrator.
- **CKAD**: Certified Kubernetes Application Developer


<hr>


## Resources:

- [Kubernetes for the absolute beginners](https://cegedim-france.udemy.com/course/learn-kubernetes/learn/lecture/9703196#overview)