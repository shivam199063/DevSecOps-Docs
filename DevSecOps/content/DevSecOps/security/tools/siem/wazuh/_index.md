---
title: "Wazuh on Kubernetes"
description: "Deploying and operating Wazuh SIEM using GitOps and Kustomize"
type: docs
toc: false
---

This section documents a **GitOps-based Wazuh deployment on Kubernetes**.

## Architecture

- Wazuh Indexer
- Wazuh Manager
- Wazuh Dashboard
- Secure inter-component communication

## Highlights

- Kustomize base + overlays
- Minikube-compatible setup
- AWS log integrations
- Kubernetes audit logs

## Deployment

🧪 **Environment**: Minikube  
🔁 **Controller**: Argo CD  

👉 GitOps application:
