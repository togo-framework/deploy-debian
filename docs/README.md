# deploy-debian — docs

**Debian VPS deploy.** SSH/systemd VPS driver for Debian hosts (apt prerequisites).

## Install

```bash
togo install togo-framework/deploy-debian
```

Registers on the [`deploy`](https://github.com/togo-framework/deploy) base; select it with **deploy.provider in togo.yaml (or DEPLOY_PROVIDER)**, then use **`togo deploy`**.

## Interface

`Deployer` — `Provision`/`Deploy`/`Destroy`/`Status` over a `Spec{App,Dir,BuildCmd,Host,User,Image,Region,Domain}` built from your `togo.yaml`.

## Usage & notes

Same SSH/systemd flow as `deploy-ubuntu` with apt package install. Config from `togo.yaml`: `host` (required), `user` (default `root`).

## Example

```bash
togo deploy --provider debian --dry-run   # preview the plan
togo deploy --provider debian
```

## Links

- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/deploy-debian)
