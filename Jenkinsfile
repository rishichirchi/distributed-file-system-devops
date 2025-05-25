pipeline {
    agent any
    environment {
        KUBECONFIG = "/home/rishi/.kube/config"
    }

    stages {
        stage('Deploy to Kubernetes') {
            steps {
                sh '''
                    kubectl apply -f ./k8s/
                '''
            }
        }
    }
}
