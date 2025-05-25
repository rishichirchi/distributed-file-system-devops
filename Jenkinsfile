pipeline {
    agent any
    environment {
        KUBECONFIG = "/home/rishi/.kube/config"
    }

    stages {
        stage('Deploy to Kubernetes') {
            steps {
                sh '''
                   sudo kubectl apply -f ./k8s/
                '''
            }
        }
    }
}
