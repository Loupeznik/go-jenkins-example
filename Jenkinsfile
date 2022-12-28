pipeline {
    agent any

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
                    sh 'export GOOS="linux"; go build -o build/app github.com/loupeznik/go-jenkins-example'
                    sh 'export GOOS="windows"; go build -o build/app.exe github.com/loupeznik/go-jenkins-example'
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
                googleStorageUpload bucket: 'gs://cclil-jenkins-tests/btc-prices-example-app', credentialsId: 'extended-argon-314215', pathPrefix: 'build/', pattern: 'build/**/*'
            }
        }
    }
}
