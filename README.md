# prometheus-stuffs

- Check diffrent branches for different project related to prometheus monitoring
    - Branch's Brief Description:
        - `prometheus-demo`:
            - just the basic server and it's monitoring by prometheus
        - `prometheus-demo-with-grafana`:
            - added prometheus operator and related things with Grafana as visualization tool with `prometheus-demo` branch codes

## How to use this branch

- `kubectl apply -f manifests/prom-rbac.yaml`
- `kubectl apply -f manifests/prom-deploy.yaml`
- `kubectl apply -f manifests/prom-server-rbac.yaml`
- `kubectl apply -f manifests/prom-server.yaml`
- `kubectl apply -f manifests/prom-server-svc.yaml`
- `kubectl apply -f manifests/grafana.yaml`
- `kubectl apply -f manifests/demo-deployment.yaml`
- `kubectl apply -f manifests/demo-svc.yaml`
- `kubectl port-forward svc/prometheus-demo-svc 8080`
- `kubectl apply -f manifests/demo-app-service-monitor.yaml`
- `kubectl port-forward svc/prometheus 9090`
- `kubectl <grafana_pod> 3000 `

- Now go to these:
    - `http://localhost:9090`: For prometheus dashboard and go to `Status` --> `targets` --> `Graphs` --> use `http_requests_total` query
    - the Visit `http://localhost:3000` for grafana
        - Our prometheus server exposing metrics at http://prometheus.default.svc:9090 (format: http://service-name.namespace.svc:port) endpoint.
        - Now we need to add data source to grafana dashboard. We need to find the prometheus service endpoint where it is exposing its data.
        - go to `+` --> `Dashboard` --> `Add Query` --> select your monitored app from the field `Query`
        - now write query in the below box to see the graph, ex: `http_requests_total` after writing this click your mouse pointer in outside of that box to apply the query.

## Basics

- Prometheus was created to monitor highly dynamic container environments, like: kubernetes
- constantly monitor all the services
- alert when a service crash
- identify problems before the event occur and alerts the system administrators (you can set the threshold)
- So, it is basically automated monitoring and alerting tools 

## How does it work

- It's core/main component is `Prometheus Server`
    - does the actual monitoring work
    - It made up of three parts
        - Time Series Database
            - stores all the metrics data, like:
                - CPU usage
                - Number of exceptions in an application
        - Data Retrieval Worker
            - it pulls/gets metrics data from:
                - applications
                - services, server,.. and other target resources
                - and store/push them into the database
        - Web sever or server api, Accepts PromQL queries
            - accepts queries for that stored data
            - server api is used to display the data in a dashboard or UI
                - like: prometheus UI or Grafana etc.

# Resources

- [ ] [Monitoring App Using Prometheus and Grafana](https://github.com/searchlight/WebApi-Prometheus/blob/master/guide/monitoring-guide.md)
- [x] [How Prometheus Monitoring works | Prometheus Architecture explained](https://www.youtube.com/watch?v=h4Sl21AKiDg)
- [ ] [Setup Prometheus Monitoring on Kubernetes using Helm and Prometheus Operator](https://www.youtube.com/watch?v=QoDqxm7ybLc)
- [ ] [Prometheus Monitoring - Steps to monitor third-party apps using Prometheus Exporter](https://www.youtube.com/watch?v=mLPg49b33sA)
- [ ] https://prometheus.io/
- [ ] https://www.youtube.com/watch?v=CmPdyvgmw-A
- [ ] https://www.youtube.com/watch?v=bErGEHf6GCc
- [ ] https://www.youtube.com/watch?v=losQlALIsYY
- [ ] https://github.com/appscode/third-party-tools
- [ ] [promql-cheat-sheet](https://promlabs.com/promql-cheat-sheet/)
- [ ] [Prometheus Example App](https://github.com/brancz/prometheus-example-app)
- [ ] https://github.com/prometheus/prometheus
