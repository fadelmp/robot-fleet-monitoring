pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'docker build -t robot-fleet-monitoring-image:latest .'
            }
        }
        stage('Push') {
            steps {
                sh 'docker push robot-fleet-monitoring-image:latest'
            }
        }
        stage('Deploy') {
            steps {
                sh 'kubectl apply -f deployment.yaml'
                sh 'kubectl apply -f service.yaml'
            }
        }
    }
}
