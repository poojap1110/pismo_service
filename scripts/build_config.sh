#!/usr/bin/env bash
# exit on error

set -eu

cd ${APP_DIR}

echo ">>>>>>>> Creating env file..."
for env in ${1}
do
    echo "$env="$(eval echo \$$env) >> ${APP_ENV_FILE}
done

echo ">>>>>>>> Creating supervisord..."
rm -rf /etc/supervisor/supervisord.conf
sh scripts/create_supervisord.sh ${APP_NAME} > supervisord.conf