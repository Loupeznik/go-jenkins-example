pipeline {
    agent any
    tools {
        go 'go'
    }
    environment {
        GO119MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Initial') {
            steps {
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin:${HOME}/go/bin"]) {
                    echo 'Installing dependencies'
                    sh 'go get .'
                }
            }
        }

        stage('Build') {
            steps {
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin:${HOME}/go/bin"]) {
                    echo 'Compiling and building'
                    sh 'go build .'
                }
            }
        }

        stage('Test') {
            steps {
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin:${HOME}/go/bin"]) {
                    echo 'Running tests'
                    sh 'go test -v'
                }
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
