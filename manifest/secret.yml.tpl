apiVersion: v1
data:
  tls.crt: ${cer}
  tls.key: ${key}
kind: Secret
metadata:
  name: frontend-tls
type: Opaque
