# Setup k8s-context
local_resource('set k8s context', 'kubectl config use-context k3d-local')

# Validate this context
if k8s_context() != 'k3d-local':
  fail("failing early as k8s context is not setup to use 'k3d-local'")

# Compile the go binary locally only when the source files are updated
local_resource(
  'Compile',
  'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o eventrouter ./',
  deps=['main.go', 'eventrouter.go', 'sinks'],
)

# Build docker image using local binary only when the main binary has changed
docker_build(
    'index.docker.io/mintel/eventrouter:0.3.0_mintel.0.1.0',
    context='.',
    dockerfile='Dockerfile.dev',
    only=['eventrouter'],
)

# Render manifests
k8s_yaml(local("tk show --dangerous-allow-redirect ./jsonnet/environments/eventrouter/local"))

# Watch for changes in manifests source
watch_file('jsonnet')
