# KubeGazer

Kubegazer is a lightweight, high-performance Kubernetes cluster observability dashboard. 

Its purpose will be to provide real-time node resource tracking, pod status aggregation, and workload diagnostics without the overhead of heavy metrics infrastructure or third-party agent dependencies.

---

## Architecture

```text
[ Browser / UI ] 
       │ (REST / JSON)
       ▼
[ pkg/api ] ── HTTP Server & Request Handlers
       │
[ pkg/k8s ] ── Domain DTOs & Resource Translators (Nodes / Pods)
       │
[ k8s.io/client-go ]
       │
       ▼
[ Kubernetes API Server ]

much more to come!    
