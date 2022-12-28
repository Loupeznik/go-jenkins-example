pipeline {
    agent any
    tools {
        go 'go1.19'
    }
    environment {
        GO119MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Initial') {
            steps {
                echo 'Installing dependencies'
                sh 'go get .'
            }
        }

        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build .'
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests'
                sh 'go test -v'
            }
        }

        stage('Publish') {
            steps {
                echo 'Publishing to artifact storage'
                googleStorageUpload bucket: 'gs://cclil-jenkins-tests', credentialsId: 'extended-argon-314215', pattern: '*.exe'
            }
        }
    }
}
