pipeline {
    agent any
    stages {
        stage ('Prepare stag env') {
            steps {
                sh "./delete-kong.sh"
                sh "docker rm -f users-service users-arangodb"
            }
        }
        stage('Build') {
            steps {
                sh "docker-compose build"
            }
        }
        stage ('Deploy') {
            steps {
                sh "docker-compose up -d --force-recreate"
                sh "./create-kong.sh"
            }
        }
    }
    post {
        always {
            sh "docker-compose down --rmi='all'"
            sh "docker system prune -f"
        }
    }
}
