pipeline {
    agent any

    environment {
        COMPOSE = "docker-compose"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh "${COMPOSE} build"
            }
        }

        stage('Up') {
            steps {
                sh "${COMPOSE} up -d"
            }
        }

        stage('Tests') {
            steps {
                sh "bash tests/test_services.sh"
            }
        }
    }

    post {
        always {
            sh "${COMPOSE} down"
        }
    }
}