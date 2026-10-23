pipeline {
    agent any

    environment {
        // Using BUILD_NUMBER ensures every push gets a unique Docker tag
        DOCKER_IMAGE = 'aadithyaar/cia2-app'
        IMAGE_TAG = "${env.BUILD_NUMBER}"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                sh "docker build -t ${DOCKER_IMAGE}:${IMAGE_TAG} -t ${DOCKER_IMAGE}:latest ."
            }
        }

        stage('Push to Docker Hub') {
            steps {
                // Ensure your Docker credentials ID in Jenkins matches this (e.g., 'docker-creds')
                withCredentials([usernamePassword(credentialsId: 'docker-creds', usernameVariable: 'DOCKER_USR', passwordVariable: 'DOCKER_PSW')]) {
                    sh "echo \$DOCKER_PSW | docker login -u \$DOCKER_USR --password-stdin"
                    sh "docker push ${DOCKER_IMAGE}:${IMAGE_TAG}"
                    sh "docker push ${DOCKER_IMAGE}:latest"
                }
            }
        }

        stage('Deploy to Azure AKS') {
            steps {
                withKubeConfig([credentialsId: 'azure-aks-kubeconfig']) {
                    // This forces the UTF-8 encoding required by Ansible before running the playbook
                    sh 'export LC_ALL=C.UTF-8 && export LANG=C.UTF-8 && ansible-playbook deploy.yml'
                }
            }
        }
    }
}