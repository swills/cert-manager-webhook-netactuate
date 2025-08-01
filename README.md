<p>
  <img src="https://raw.githubusercontent.com/cert-manager/cert-manager/d53c0b9270f8cd90d908460d69502694e1838f5f/logo/logo-small.png" height="256" width="256" alt="cert-manager project logo" />
</p>

# ACME webhook for netactuate

[![Go](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/build.yml/badge.svg)](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/build.yml)
[![golangci-lint](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/golangci-lint.yml)
[![pages-build-deployment](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/pages/pages-build-deployment/badge.svg)](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/pages/pages-build-deployment)
[![Release Image](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/release-docker.yml/badge.svg)](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/release-docker.yml)
[![Release Charts](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/release-charts.yml/badge.svg)](https://github.com/swills/cert-manager-webhook-netactuate/actions/workflows/release-charts.yml)

## How to use the helm chart:

Assuming you already have cert-manager deployed in the cert-manager namespace using helm:

```bash
helm repo add swills-cert-manager-webhook-netactuate https://swills.github.io/cert-manager-webhook-netactuate/
```

Ensure your IP(s) are allowed in the IP ACLs, see [this note](https://status.netactuate.com/pages/maintenance/59cd452c6a99786b77cab2f7/67d1c5fe63ec070537f4b651)


Create your api key secret:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: netactuate-api-key
  namespace: cert-manager
data:
  netactuate-api-key: cmVwbGFjZS13aXRoLW5ldGFjdHVhdGUtYXBpLWtleQ==
```

Create your cluster issuer:
```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: acme-example-com
spec:
  acme:
    email: you@example.com
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: acme-example-com-clusterissuer-private-key-secret
    solvers:
      - dns01:
          webhook:
            config:
              apiKey:
                name: netactuate-api-key
                value: netactuate-api-key
            groupName: acme.example.com
            solverName: netactuate
        selector:
          dnsZones:
            - example.com
```

Deploy the chart:
```bash
helm install --namespace cert-manager netactuate-webhook swills-cert-manager-webhook-netactuate/netactuate-webhook
```

## How to test
```bash
$ env NETACTUATE_API_KEY='your-api-key' TEST_DOMAIN="example.coM." TEST_RECORD_ID=123456 go test -v ./...
```

Note: You must change these example values to match your account API key, test domain and test record ID.
