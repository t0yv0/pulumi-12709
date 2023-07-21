#!/usr/bin/env bash

set -euo pipefail

pulumi stack select $(pulumi stack ls -j | jq -r '.[0].name')
export HC_SKIP_VPN_CHECK=true
export PATH=$PWD/../../bin:$PATH
export PULUMI_CONFIG_PASSPHRASE=1234567

pulumi config set region us-west-2
pulumi destroy --yes

rm -rf "$PWD/log.json"
rm -rf "$PWD/log-formatted.json"
PULUMI_DEBUG_GRPC="$PWD/log.json" pulumi up --yes || echo ignoreFAIL
jq . "$PWD/log.json" > "$PWD/log-formatted.json"
