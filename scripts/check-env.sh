. ./common.sh

EXECUTABLES="go solc docker terraform jq"

echo "Checking scripts dependencies:"

lookFor() {
  echo "\tlooking for '$1' ..."
  command -v $1 >/dev/null 2>&1
  if [ "$?" != "0" ]; then
      fail "\t\t'$1' not found!"
  fi
}

for e in $EXECUTABLES; do lookFor $e; done

bye "All required tools are installed"

