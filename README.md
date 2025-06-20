# 🛠️ Kuma Diagnoser CLI

Una herramienta sencilla escrita en Go para diagnosticar el estado del service mesh Kuma, enfocándose en dataplanes y control plane.

## Comandos

### `check`
Realiza una verificación rápida del estado del control plane y dataplanes.

```bash
kuma-diagnoser check
kuma-diagnoser export --format md --output diag.md
kuma-diagnoser port-forward
