pipeline {

    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
    }
    stages {
        stage ("Build") {
            steps {
                sh "docker-compose build --pull"
            }
        }
        stage ("Deploy") {
            steps {
                sh "docker-compose up -d --force-recreate"
            }
        }
    }
    post {
        always {
            sh "docker system prune -f"
        }
    }
}