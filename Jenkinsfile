pipeline {
    agent any
    stages {
        stage ('Deploy') {
            steps {
                catchError {
                    sh "docker-compose down"
                    sh "docker-compose up -d --build --force-recreate"
                    sh "./create-kong.sh"
                }
            }
        }
    }
    post {
        always {
            sh "docker system prune -f"
        }
        failure {
            sh "./delete-kong.sh"
            sh "docker rm -f users-service users-arangodb"
        }
    }
}
