# Summary

The platform bootstrapper is a project that is responsible for the following items:

- Creating the initial namespace where the platform operators exist
- Deploying the [platform-config-operator](https://github.com/tbd-paas/platform-config-operator)
- Deploying the [platform capability operator definitions](https://github.com/tbd-paas/platform-config-operator/blob/main/config/crd/bases/deploy.platform.tbd.io_platformoperators.yaml)
- Ensuring the platform capability operator deployments are successful
- Deploying the [platform capability config definitions](https://github.com/tbd-paas/platform-config-operator/blob/main/config/crd/bases/deploy.platform.tbd.io_platformconfigs.yaml)
- Ensuring the platform capability components are successfully deployed
- Streaming logs via http so that any process integrating with the bootstrapper service may retrieve the latest status

