# MySQLRoute

- https://github.com/voyagermesh/gateway/blob/main/docs/v0.4.0/user/tcp-routing.md
- https://github.com/envoyproxy/envoy/tree/main/examples/mysql
- https://organic-alder-bf8.notion.site/envoy-gateway-a21aa3a2c5fd4448b865eae41420afa3

```yml
cat <<EOF | kubectl apply -f -
kind: GatewayClass
apiVersion: gateway.networking.k8s.io/v1beta1
metadata:
  name: eg
spec:
  controllerName: gateway.envoyproxy.io/gatewayclass-controller
---
apiVersion: gateway.networking.k8s.io/v1beta1
kind: Gateway
metadata:
  name: db-gateway
spec:
  gatewayClassName: eg
  listeners:
  - name: foo
    protocol: TCP
    port: 3306
    allowedRoutes:
      kinds:
      - group: gateway.voyagermesh.com
        kind: MySQLRoute
  - name: bar
    protocol: TCP
    port: 27017
    allowedRoutes:
      kinds:
      - group: gateway.voyagermesh.com
        kind: MongoDBRoute
EOF
```

```yml
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  name: foo
  labels:
    app: backend-1
spec:
  ports:
    - name: http
      port: 3001
      targetPort: 3000
  selector:
    app: backend-1
---
apiVersion: v1
kind: Service
metadata:
  name: bar
  labels:
    app: backend-2
spec:
  ports:
    - name: http
      port: 3002
      targetPort: 3000
  selector:
    app: backend-2
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend-1
spec:
  replicas: 1
  selector:
    matchLabels:
      app: backend-1
      version: v1
  template:
    metadata:
      labels:
        app: backend-1
        version: v1
    spec:
      containers:
        - image: gcr.io/k8s-staging-ingressconformance/echoserver:v20221109-7ee2f3e
          imagePullPolicy: IfNotPresent
          name: backend-1
          ports:
            - containerPort: 3000
          env:
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
            - name: SERVICE_NAME
              value: foo
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend-2
spec:
  replicas: 1
  selector:
    matchLabels:
      app: backend-2
      version: v1
  template:
    metadata:
      labels:
        app: backend-2
        version: v1
    spec:
      containers:
        - image: gcr.io/k8s-staging-ingressconformance/echoserver:v20221109-7ee2f3e
          imagePullPolicy: IfNotPresent
          name: backend-2
          ports:
            - containerPort: 3000
          env:
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
            - name: SERVICE_NAME
              value: bar
EOF
```

```yml
cat <<EOF | kubectl apply -f -
apiVersion: gateway.voyagermesh.com/v1alpha1
kind: MySQLRoute
metadata:
  name: tcp-app-1
spec:
  parentRefs:
  - name: db-gateway
    sectionName: foo
  rules:
  - backendRefs:
    - name: foo
      port: 3001
---
apiVersion: gateway.networking.k8s.io/v1alpha2
kind: MongoDBRoute
metadata:
  name: tcp-app-2
spec:
  parentRefs:
  - name: db-gateway
    sectionName: bar
  rules:
  - backendRefs:
    - name: bar
      port: 3002
EOF
```
