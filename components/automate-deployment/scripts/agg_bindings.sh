#!/usr/bin/env bash

if [ "${BASH_VERSINFO[0]}" -lt 4 ]
then
  echo "Minimum version of bash 4 is required. Found ${BASH_VERSINFO[0]}."
  echo "Try running inside the hab studio instead."
  exit 1
fi

BINDS_FILE=./pkg/assets/data/binds.txt

bindings() {
  for component_path in $(find ../../components/* -maxdepth 0 -type d | sort)
  do
    # We'll check anything with a runhook
    if [[ -e ${component_path}/habitat/hooks/run ]]
    then
      (
        declare -A pkg_binds
        declare -A pkg_binds_optional
        local chef_automate_hab_binding_mode="strict"
        local chef_automate_dev_only_pkg=false
        #shellcheck disable=SC1090
        source "${component_path}/habitat/plan.sh"; 

        if [[ "${chef_automate_dev_only_pkg}" == "true" ]]; then
          # skip
          exit 0
        fi

        if [[ ! -z  ${!pkg_binds[@]} ]]
        then
          readarray -t sorted < <(printf '%s\n' "${!pkg_binds[@]}" | sort)
          #shellcheck disable=SC2154
          echo "${pkg_name}" REQUIRED "${sorted[@]}"
        fi

        if [[ ! -z  ${!pkg_binds_optional[@]} ]]
        then
          readarray -t sorted < <(printf '%s\n' "${!pkg_binds_optional[@]}" | sort)
          echo "${pkg_name}" OPTIONAL "${sorted[@]}"
        fi
        echo "${pkg_name}" BINDING_MODE "${chef_automate_hab_binding_mode}"
      ) 2> /dev/null
    fi
  done
}

dryrun='false'
for i in "$@"
do
  case $i in
    -n)
      dryrun='true'
      shift
      ;;
    *)
      ;;
  esac
done

if [[ $dryrun == 'true' ]]
then
  diff $BINDS_FILE <(bindings) > /dev/null
  if [[ $? != 0 ]]
  then
    cat <<EOF
It looks like the bindings for a service have changed. The binds.txt file in
the deployment service must be updated. To fix this issue, go to
components/automate-deployment and run:

  > make update-bindings

EOF

  exit 1
  fi
else
  if ! $(bindings | sort |sort -uc); then
    cat <<EOF
Detected multiple packages with duplicate pkg_name set in plan.sh
Dev-only packages should not be used for bindings data. To mark a package to be
ignored for binding data, add the following to the plan.sh:

  chef_automate_dev_only_pkg=true
EOF
    exit 1
  fi

  # If a data service can be operated in an external mode, we do not want other
  # services to bind to it. Currently this is just elastic.
  # * automate-es-gateway is allowed to bind to it because it is the proxy that
  # sends requests to the internal elastic when using internal elastic
  # * es-sidecar is allowed to bind to elastic for now because our initial
  # implementation of external elastic requires it for backups to function.
  services_binding_to_elastic=$(bindings | grep -E 'REQUIRED|OPTIONAL' | grep -v ^automate-elasticsearch | grep -v ^automate-es-gateway | grep -v ^es-sidecar-service | grep automate-elasticsearch)
  if [[ "${services_binding_to_elastic}" != "" ]]; then
    cat <<EOF
Detected packages with bindings on elasticsearch:

${services_binding_to_elastic}

Elasticsearch may be deployed external to automate and may therefore be
unavailable. Services that use elasticsearch should access it via
automate-es-gateway.
EOF
    exit 1
  fi

  :> $BINDS_FILE
  bindings >> $BINDS_FILE
  exit 0
fi
