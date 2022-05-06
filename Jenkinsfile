pipeline {
    agent any

    stages {
        stage('Get code') {
            steps {
                checkout scm
            }
        }
		stage('Build and deploy') {
            steps {
                withCredentials([string(credentialsId: 'T01_HOST_NAME', variable: 'NAME'), string(credentialsId: 'T01_BACKEND_HOST', variable: 'HOST'), usernamePassword(credentialsId: 'T01_HOST_ACCOUNT', passwordVariable: 'PASSWORD', usernameVariable: 'USER'), string(credentialsId: 'T01_TOKEN_KEY', variable: 'TOKEN_KEY'), string(credentialsId: 'T01_DB_HOST', variable: 'DB_HOST'), string(credentialsId: 'T01_DB_USER', variable: 'DB_USER'), string(credentialsId: 'T01_DB_PASSWORD', variable: 'DB_PASSWORD'), string(credentialsId: 'T01_DB_NAME', variable: 'DB_NAME')]) {
                    script {
                        def remote = [:]
                        remote.name = NAME
                        remote.host = HOST
                        remote.user = USER
                        remote.password = PASSWORD
                        remote.allowAnyHosts = true
                        try {
							sshCommand remote: remote, command: "git clone -b main https://github.com/up1/workshop-continuous-testing.git"
            
                            // Build & Run Container
                            sshCommand remote: remote, command: "cd workshop-continuous-testing && sh deploy.sh"

                        }  catch (err) {
                            echo "Error: ${err.message}"
                            echo "Error: ${err.stack}"
                        } finally {
                            // Clear files
                            sshRemove remote: remote, path: "workshop-continuous-testing"
                        }
                    }
                }
            }
        }
		stage('API Testing') {
            steps {
                sh 'cd testing && newman run somkiat.postman_collection.json'
            }
        }
		stage('Notify result') {
            steps {
                echo 'TODO'
            }
        }
    }
}
