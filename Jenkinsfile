pipeline {
    agent any
    environment {
        KUBECONFIG = "/home/rishi/.kube/config"
    }

    stages {
        stage('Deploy to Kubernetes') {
            steps {
                sh '''
                    minikube update-context
                    kubectl config use-context minikube
                    kubectl apply -f ./k8s/
                '''
            }
        }
    }
}
