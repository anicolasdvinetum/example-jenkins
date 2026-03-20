pipeline {
    agent any

    environment {
        COMPOSE = "docker compose"
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

        stage('Deploy containers') {
            steps {
                sh "${COMPOSE} up -d"
            }
        }

        stage('Run tests') {
            steps {
                sh "bash tests/test_services.sh"
            }
        }
    }

    post {
        always {
            echo "Cleaning up containers..."
            sh "${COMPOSE} down"
        }
    }
}