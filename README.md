# prometheus-stuffs

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

## How to use this branch's code

- This branch code is basically creating a basic API server and build a docker image of it and then uploaded it in docker hub, then used that image in k8s by deployment.yaml
- `kubectl apply manifests/deployment.yaml`: To run the image in k8s inside a pod
- now need to por forward for accessing the api endpoints from the local machine, by:
    - `kubectl port-forward <pod_name> 8080`
- now the endpoints are:
    - `curl http://localhost:8080/`: server endpoint
    - `curl http://localhost:8080/err`: server endpoint
    - `curl http://localhost:8080/metrics`: to see the server metrics through prometheus


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
