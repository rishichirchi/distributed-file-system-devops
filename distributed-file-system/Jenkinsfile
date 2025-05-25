pipeline {
    agent any

    stages {
        stage('Start Minikube') {
            steps {
                sh 'minikube start'
            }
        }

        // stage('Build') {
        //     steps {
        //         sh 'docker build -t rishichirchi/api-server ./api-server'
        //         sh 'docker build -t rishichirchi/control-plane ./control-plane'
        //         sh 'docker build -t rishichirchi/storage-node-1 ./storage-node/node1'
        //         sh 'docker build -t rishichirchi/storage-node-2 ./storage-node/node2'
        //         sh 'docker build -t rishichirchi/storage-node-3 ./storage-node/node3'
        //     }
        // }

        // stage('Push to DockerHub') {
        //     steps {
        //         withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'rishichirchi', passwordVariable: 'Do@rishi04')]) {
        //             sh 'echo $PASSWORD | docker login -u $USERNAME --password-stdin'
        //             sh 'docker push rishichirchi/api-server'
        //             sh 'docker push rishichirchi/control-plane'
        //             sh 'docker push rishichirchi/storage-node-1'
        //             sh 'docker push rishichirchi/storage-node-2'
        //             sh 'docker push rishichirchi/storage-node-3'
        //         }
        //     }
        // }

        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl apply -f ./k8s/'
            }
        }
    }
}
