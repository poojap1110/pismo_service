cd ${APP_DIR}

echo ">>>>>switch context to ${KOPS_CLUSTER_NAME}"
kubectl config use-context ${KOPS_CLUSTER_NAME}

echo ">>>>>check if ${APP_NAME} already exists"
kubectl get deployment ${APP_NAME} >> checkdeployment.txt

cat checkdeployment.txt

if [ -s checkdeployment.txt ]
then
    echo ">>>>>update docker image to ${ECR_APP_REGISTRY}/${APP_NAME}:${CI_COMMIT_ID}"
    kubectl set image deployments/${APP_NAME} ${APP_NAME}=${ECR_APP_REGISTRY}/${APP_NAME}:${CI_COMMIT_ID}

    echo ">>>>>updating secret credentials"
    kubectl create secret generic ${APP_NAME}-credentials --save-config --from-env-file=${APP_ENV_FILE} --dry-run -o yaml | kubectl apply -f -

else
    echo ">>>>>preparing the deployment yaml file"
    sh scripts/create_k8s_deployment.sh ${APP_NAME} ${ECR_APP_REGISTRY} ${CI_COMMIT_ID} ${APP_PORT} ${CONTAINER_LIMITS_CPU} ${CONTAINER_LIMITS_MEMORY} ${CONTAINER_REQUESTS_CPU} ${CONTAINER_REQUESTS_MEMORY} ${PODS_MIN_REPLICAS} > k8s-deployment.yml

    echo ">>>>>creating secret credentials"
    kubectl create secret generic ${APP_NAME}-credentials --from-env-file=${APP_ENV_FILE}

    echo ">>>>>printing deployment yaml file"
    cat k8s-deployment.yml

    echo ">>>>>creating the deployment"
    kubectl create -f k8s-deployment.yml

    echo ">>>>>setting up the horizontal pod autoscaling"
    kubectl autoscale deployments ${APP_NAME} --min=${PODS_MIN_REPLICAS} --max=${PODS_MAX_REPLICAS} --cpu-percent=${PODS_TARGET_CPU_UTILIZATION_PERCENTAGE}

    echo ">>>>>creating service discovery yaml file"
    for service in $(echo ${SERVICE_DISCOVERIES} | sed "s/,/ /g")
    do
        sh scripts/create_k8s_service_discovery.sh ${APP_NAME} ${SERVICE_PORT} ${APP_PORT} ${service}  >> k8s-services.yml
    done

    echo ">>>>>printing service yaml file"
    cat k8s-services.yml

    echo ">>>>>creating the service(s) discovery"
    kubectl create -f k8s-services.yml

fi