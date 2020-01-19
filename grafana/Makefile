SHELL=/bin/bash -euo pipefail
diff:
	go run . | jsonnet -J vendor/ - | jq . > out.json
	(diff -u -N data.json out.json | colordiff || true) | less
