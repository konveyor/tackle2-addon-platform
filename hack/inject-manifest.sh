#!/bin/bash

while getopts ":v:m:h" opt; do
  case $opt in
    v) values="$OPTARG"
    ;;
    m) manifest="$OPTARG"
    ;;
    h)
      echo "Usage: $0 [options]"
      echo "Options:"
      echo "  -v <path>  Path to values file (default: values.yaml)"
      echo "  -m <path>  Path to manifest file (default: manifest.yaml)"
      echo "  -h         Display this help message"
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


