# GitLab CI Docker Credentials Helper

This credentials helper is meant for use in GitLab CI environments. It serves the container registry at `CI_REGISTRY` with the credentials from `CI_REGISTRY_USER` and `CI_REGISTRY_PASSWORD`, and the dependency proxy at `CI_DEPENDENCY_PROXY_SERVER` with the credentials from `CI_DEPENDENCY_PROXY_USER` and `CI_DEPENDENCY_PROXY_PASSWORD`.
