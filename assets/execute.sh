#!/bin/sh
set -eu
_TMP_DIR="$(mktemp -d)"
echo "{{ .TarBallBase64 }}" | base64 -d >"$_TMP_DIR/tarball"
mkdir -p "$_TMP_DIR/extract"
tar -xf "$_TMP_DIR/tarball" -C "$_TMP_DIR/extract"
{
    echo "#!{{ .Shebang }}" >"$_TMP_DIR/execute"
    echo "source $_TMP_DIR/extract/{{ .Entrypoint }}"
    echo "export SEIRA_ROOTDIR=\"$_TMP_DIR/extract\""
} >>"$_TMP_DIR/execute"
echo 'main "$@"' >>"$_TMP_DIR/execute"
chmod +x "$_TMP_DIR/execute"
"$_TMP_DIR/execute"
