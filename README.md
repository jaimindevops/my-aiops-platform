# 🧠 AIOps-K8s-Platform: Self-Healing & AI-Powered DevSecOps

**A production-grade Kubernetes platform built on Oracle Cloud (OCI) that leverages Local AI for observability, Chaos Engineering for resilience, and GitOps for automated security.**

---

## 📖 Executive Summary

As modern infrastructure grows complex, "keeping the lights on" requires more than just monitoring—it requires intelligence. This project demonstrates a **Self-Healing AIOps Platform** that detects issues using Generative AI, blocks security threats automatically, and survives catastrophic failures without data loss.

### 🎯 Key Achievements

* **94% Attack Surface Reduction:** Migrated Go application to **Google Distroless** images (246MB → 14MB).
* **Zero-Downtime Resilience:** Implemented **Redis AOF Persistence** with PVCs; data survived forced pod deletion during chaos testing.
* **AI-Driven Troubleshooting:** Integrated **k8sgpt** with a local LLM (**Gemma 2B**) to analyze crash loops and explain errors in plain English without sending data to the cloud.
* **Automated Security Governance:** configured GitHub Actions to block builds when **Trivy** detects critical CVEs.

---

## 🏗️ Architecture
<img width="809" height="459" alt="image" src="https://github.com/user-attachments/assets/ce92ed92-18b7-4ab0-8a14-65c6851a8fb8" />


The platform runs on **Oracle Cloud Infrastructure (OCI)** using the ARM64 Ampere tier for maximum efficiency.

| Component | Technology Stack |
| --- | --- |
| **Orchestration** | **K3s** (Lightweight Kubernetes) |
| **Ingress Controller** | **Traefik** (L7 Load Balancing & Routing) |
| **Database** | **Redis** (StatefulSet with AOF Persistence) |
| **Observability (AI)** | **k8sgpt** + **Ollama** (Local LLM: Gemma 2B) |
| **Security Scanning** | **Trivy** (CI/CD Integration) |
|**Container Engine** | **containerd** (Managed by K3s) |
|**Secure Base Image** | **Google Distroless** (Static Debian 12) |
|**Build Tool** | **Docker** (Multi-Stage builds) |
| **Chaos Engineering** | **Chaos Mesh** (Fault Injection) |
| **CI/CD** | **GitHub Actions** + **ArgoCD** (GitOps) |

---

## 🚀 Key Features & Evidence

### 1. AIOps: "The Intelligent Operator"

Instead of manually digging through `kubectl logs`, the platform uses an embedded AI agent to diagnose issues.

* **Tooling:** `k8sgpt` connected to a local `Ollama` instance running `gemma:2b`.
* **Workflow:** The AI detects a `CrashLoopBackOff`, analyzes the error code, and suggests a fix (e.g., "ImagePullBackOff detected: Check registry credentials").
* **Privacy:** 100% Local. No cluster data is sent to OpenAI or external APIs.

### 2. DevSecOps: "The Ironclad Pipeline"

Security is shifted left. The CI/CD pipeline enforces a **hard gate** on vulnerabilities.

* **Optimization:**
* *Before:* `golang:alpine` (246MB)
* *After:* `gcr.io/distroless/static` (14MB)
* **Result:** **94% reduction** in size and attack surface.


* **Vulnerability Blocking:**
* Trivy scans the image during the build process.
* If `CRITICAL` vulnerabilities are found, the pipeline **fails** preventing deployment.
* *See `screenshots/trivy-failure.png` for the security gate in action.*
<img width="839" height="368" alt="image" src="https://github.com/user-attachments/assets/c6d8b1af-9028-4647-8e3d-a02202a132b2" />


### 3. Stateful Resilience: "The Chaos Test"

Stateless apps are easy; stateful apps are hard. I built a visitor counter to prove data persistence.

* **Scenario:** A visitor counter app connected to Redis.
* **The Test:** Manually deleted the Redis pod (`kubectl delete pod redis-0`) while the app was live.
* **The Result:** Kubernetes spun up a new pod, re-attached the **Persistent Volume Claim (PVC)**, and the visitor count resumed from **9** instead of resetting to **1**.

---

## 📸 Screenshots & Proof of Work

### ✅ Evidence 1: Data Persistence (Visitor Count = 6)

*Proof that the database survived a pod crash.*

### ✅ Evidence 2: Image Optimization (94% Reduction)

*Comparison between the original Alpine image and the hardened Distroless image.*
<img width="1057" height="177" alt="image" src="https://github.com/user-attachments/assets/b392c145-3f15-4a2e-9517-aa3710562b39" />


### ✅ Evidence 3: AIOps Analysis

*k8sgpt analyzing a broken pod and suggesting a fix.*
<img width="507" height="335" alt="image" src="https://github.com/user-attachments/assets/aac6aba1-d69f-45a7-b186-0dbf5569d564" />


### ✅ Evidence 4: Self-Healing in Action
Kubernetes automatically detecting a failed container (simulated via Chaos Mesh) and restarting it immediately to maintain availability.
<img width="533" height="111" alt="image" src="https://github.com/user-attachments/assets/b1de30a4-e561-4058-a8df-6d24bb27080a" />

---

## 🛠️ How to Replicate (Local Setup)

Want to run this yourself?

**1. Clone the Repository**

```bash
git clone https://github.com/jaiminops18/my-aiops-platform.git
cd my-aiops-platform

```

**2. Deploy to Kubernetes**

```bash
# Apply the Redis Database (Stateful)
kubectl apply -f manifests/redis-pvc.yaml
kubectl apply -f manifests/redis.yaml

# Apply the Go Application
kubectl apply -f manifests/go-app.yaml

# Apply Ingress Routes
kubectl apply -f manifests/ingress.yaml

```

**3. Configure Local DNS**
Add the following to your `/etc/hosts` file:

```bash
# Replace with your LoadBalancer IP
192.168.x.x  aiops.local

```

**4. Access the Platform**
Open your browser and navigate to: `http://aiops.local`
<img width="748" height="170" alt="image" src="https://github.com/user-attachments/assets/33622a49-92d1-49dc-b1a6-65395d7de53d" />

---

## 🔮 Future Roadmap: Agentic AI

Currently, the system uses **AIOps** (Human-in-the-loop). The next phase involves upgrading to **Agentic AI** (Human-on-the-loop), where the AI agent is granted write access to the Kubernetes API to:

1. Detect a crash.
2. Auto-revert the deployment.
3. Notify the engineering team via Slack.

---

## 👨‍💻 Author

**Jaimin Sharma**
* [LinkedIn](https://www.linkedin.com/in/jaimin-sharma/)
* [GitHub](htttps://www.github.com/jaimindevops)

---

*Built with ❤️ on Oracle Cloud Free Tier.*
