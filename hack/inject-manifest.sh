#!/bin/bash

while getopts ":v:m:h" opt; do
  case $opt in
    v) values="$OPTARG"
    ;;
    m) manifest="$OPTARG"
    ;;
    h)
      echo
      echo "Inject an application manifest into the values file under the 'manifest:' node."
      echo
      echo "usage: $0 [options]"
      echo "options:"
      echo "  -v <path>  path to values file (default: values.yaml)"
      echo "  -m <path>  path to manifest file (default: manifest.yaml)"
      echo "  -h         display help"
      exit 0
    ;;
    ?) echo "Invalid option -$OPTARG"; exit 1
    ;;
  esac
done

values="${values:-values.yaml}"
manifest="${manifest:-manifest.yaml}"

yq e -n '.manifest = load("'"${manifest}"'")' | \
  yq eval-all -i 'select(fileIndex == 0) * select(fileIndex == 1)' "${values}" -


