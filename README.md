# GitLab GoProxy

This project contains an implementation of the [`GOPROXY` protocol](https://go.dev/ref/mod#goproxy-protocol) that stores data in the [GitLab Generic Package Repository](https://docs.gitlab.com/ee/user/packages/generic_packages/).

## Getting started

1. In GitLab, create a project that you intend on using to store packages.
    > We recommend creating a new project otherwise it will make it hard to find any pre-existing packages.
2. Create a [project](https://docs.gitlab.com/ee/user/project/settings/project_access_tokens.html) or group access token with the `api` scope.
3. Create a Kubernetes secret using the token we just created:
    ```shell
    export GOPROXY_TOKEN=<token we just got>
    kubectl create secret generic goproxy-token --from-literal=token=$GOPROXY_TOKEN
    ```
4. Create a `values.yaml` file with at least the following:
    ```yaml
    gitlab:
      url: https://gitlab.example.com
      projectId: 123 # numeric ID of the project we created
      tokenSecretRef:
        name: goproxy-token
        key: token
    ```
    There's plenty more that can be figured, but we will leave that as an exercise to the reader.
5. Deploy the application:
    ```shell
    helm install goproxy ./chart/gitlab-goproxy -f values.yaml
    ```
   > We recommend using a GitOps tool (e.g., FluxCD, ArgoCD) rather than doing this manually.
   
## Usage

Once the application has been deployed, using it is as simple as setting the `GOPROXY` environment variable:

```shell
export GOPROXY="https://goproxy.example.org"
go mod download
```
