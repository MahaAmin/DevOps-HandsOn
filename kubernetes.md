# Kubernetes

## Table of Content:

1. [Kubernetes For The Absolute Beginners](#kubernetes-for-the-absolute-beginners)
    1. [Kubernetes Overview](#kubernetes-overview)
    2. [Docker Containers](#docker-containers)
    3. [Containers Orchestration](#containers-orchestration)
    4. [Kubernetes Architecture](#kubernetes-architecture)
    4. [Kubernetes Setup](#kubernetes-setup)
    6. [Minikube](#minikube)
    5. [Kubernetes Concepts (PODs - ReplicaSets - Deployment - Services)](#kubernetes-concepts)
    6. [Networking in Kubernetes](#networking-in-kubernetes)
    7. [KubeCtl](#kubectl)
    8. [YAML](#yaml)
    9. [Kubernetes on Cloud](#kubernetes-on-cloud)
2. [Kubernetes Certificates](#kubernetes-certificates)
3. [Notes](#notes)
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


### Kubernetes Setup:

- There are Multi-node and Single-node setups.

- Options to install kubernetes:
    - Install kubernetes locally on your machine.
    - MicroK8s.
    - Minikube.
    - Kubeadm.
    - On cloud provider (AWS, GCP, Microsoft-Azure)

- Minikube:
    - Minikube executable.
    - Hypervisor (VirtualBox)
    - KubeCtl.

<hr>


### Minikube:

- Start minikube:
    ```
    minikube start --driver=virtualbox
    ```

- Get minikube status:
    ```
    minikube status
    ```

- 


<hr>

### Kubernetes Concepts:

- **PODs:**
    - Smallest object you can create on kubernetes.
    - One-to-one relationship with containers : each pod contains one container.
    - To scale up --> create new pods, not adding new containers to the same pod.
    - To scale down --> delete pods.
    - **Multi-Container PODS:**
        - A single pod can have multiple containers inside it.
        - Multiple container but in **different types**.
        - They are called **Helper containers**. Containers that are highly assciociated with the main container in the pod.
        - Multiple containers are not intended for scaling, but rather for functionality.
        - Containers in the same pod shares the same network-space, and storage-space. 
        - They can communicate with eachother as "localhost", as they are in the same pod.
        - When scaling up all containers in the same pods get replicated, and in scaling down all containers die together.

<hr>

### KubeCtl:

- Kubernetes Command-line tool.
- Kube-Control.
- Kube C T L.
- Used to deploy and manage applications in kubernetes cluster.
- Get cluster's information.

- Start minikube:
    ```
    minikube start --driver=virtualbox
    ```

- Deploy an application on the cluster:
    ```
    kubectl run hello-minikube
    ```

- Get information about the cluster:
    ```
    kubectl cluster-info
    ```

- List nodes in a cluster:
    ```
    kubectl get nodes
    ```

- Get more info about the nodes:
    ```
    kubectl get nodes -o wide
    ```

- Deploy a pod from image in DockerHub:
    ```
    kubectl run nginx --image nginx
    ```

- List existing pods:
    ```
    kubectl get pods
    ```

    ```
    kubectl get pods -o wide
    ```

- Get more information about specific pod:
    ```
    kubectl describe pod nginx
    ```



<hr>


## Kubernetes Certificates:

- **CKA** : Certified Kubernetes Adminstrator.
- **CKAD**: Certified Kubernetes Application Developer


<hr>


## Notes:
- We use docker-compose for local environment, and kubernetes for production.


## Resources:

- [Kubernetes for the absolute beginners](https://cegedim-france.udemy.com/course/learn-kubernetes/learn/lecture/9703196#overview)