pipeline {
    agent any
    environment {
        KUBECONFIG = "/var/lib/jenkins/.kube/config"
    }
    stages {
        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl apply -f ./k8s/ --validate=false'
            }
        }
    }
}
