# Vanity GO on Kubernetes

The solution I personally use in my Kubernetes cluster.

In this example I will use Istio as IngressController, being the one I personally use. Of course vanity-go is compatible with any type of ingress, being a proxy at layer 7.

**You can use the example as an initial configuration.**

## Workflow

First you need to apply the workload to your cluster, once the workload is installed, you can edit the configmap to configure the vanity-go for your domain names and repositories. 

```sh
# Apply the example deployment files (configmap, deployment, service)
kubectl apply -f https://raw.githubusercontent.com/42Atomys/vanity-go/1.0/examples/kubernetes/deployment.yaml

# Edit the configuration map to apply your redirection and configurations
kubectl edit configmap/vanity-go-config
```

Don't forget to restart your deployment so that your vanity-go takes into account the changes made to your configmap
```sh
# Restart your proxy to apply the latest configuration
kubectl rollout restart deployment.apps/vanity-go
```

It's all over! 🎉

Now it depends on your Ingress!

> **REMEMBER** Remote Import Paths with always try to get package
> with query params `?go-get=1`. It's very nice if you have a homepage on your domain and don't want to add meta on this page.
> 
> I strongly recommend that you do your routing to follow this guideline. The go-get protocol will first try to return the package in HTTPS before fallback in HTTP.


## Sugar Free: Isito Routing

If you use istio as IngressController like me, you can my virtual service (it's free)

```yaml
---
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: vanity-go
spec:
  hosts:
  - atomys.codes # Change for your domain
  gateways:
  - default
  http:
  - match:
    # Remember ! Remote Import Paths with always try to get package
    # with query params ?go-get=1 before without params.
    # I recommand you to route only on this case
    - queryParams:
        go-get:
          exact: '1'
    route:
    - destination:
        port:
          number: 8080
        host: vanity-go

```