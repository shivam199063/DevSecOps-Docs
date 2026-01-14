---
title: "Connection"
description: "How to make a Kubernetes connection."
type: docs
toc: false
weight: 10
---

## Overview

This document explains **how a client connects to a Kubernetes cluster**, both from a **local machine** and from **inside the cluster**.  

---

## How Kubernetes Connections Work

All communication with a Kubernetes cluster happens through the **API Server**.


The client never talks directly to:
- kubelet
- scheduler
- controller-manager

Everything goes through the **API Server**.

---

## Connecting to Kubernetes from Local Machine

### 1. kubeconfig file

Kubernetes uses a configuration file called **kubeconfig**.

**Default location:**
```bash
~/.kube/config
```

**This file contains:**
- Cluster endpoint
- Certificate authority
- Authentication credentials
- Contexts (cluster + user + namespace)

```
apiVersion: v1
clusters:
- cluster:
    certificate-authority: /Users/shivamsaini/.minikube/ca.crt
    extensions:
    - extension:
        last-update: Sun, 11 Jan 2026 13:44:47 IST
        provider: minikube.sigs.k8s.io
        version: v1.37.0
      name: cluster_info
    server: https://127.0.0.1:49981
  name: shivam-k8s
contexts:
- context:
    cluster: shivam-k8s
    extensions:
    - extension:
        last-update: Sun, 11 Jan 2026 13:44:47 IST
        provider: minikube.sigs.k8s.io
        version: v1.37.0
      name: context_info
    namespace: default
    user: shivam-k8s
  name: shivam-k8s
current-context: shivam-k8s
kind: Config
preferences: {}
users:
- name: shivam-k8s
  user:
    client-certificate: /Users/shivamsaini/.minikube/profiles/shivam-k8s/client.crt
    client-key: /Users/shivamsaini/.minikube/profiles/shivam-k8s/client.key
```


**Verify cluster connection**
```bash
kubectl get nodes
```

### 2. Connecting Using Service Account (Inside Cluster)

- Applications running inside Kubernetes do not use kubeconfig files.

- Instead, Kubernetes automatically mounts credentials into the Pod.

What gets mounted automatically
```text
/var/run/secrets/kubernetes.io/serviceaccount/
├── token
├── ca.crt
└── namespace
```
These are used to authenticate with the API server.

### Example: In-cluster connection (concept)

- **API Server URL:** `https://kubernetes.default.svc`
- **Authentication:** ServiceAccount token
- **Authorization:** RBAC rules
