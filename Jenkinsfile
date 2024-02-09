pipeline {
    agent any

    environment {
    VERSION = "0.0.3-beta-${BUILD_NUMBER}"
    }

    stages {

        stage('Init') {
            steps {
                sh 'printenv'
            }
        }

    }

    post {
        cleanup {
            sh 'printf cleanup'
        }
    }
}