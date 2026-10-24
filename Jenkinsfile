pipeline {
    agent any
    
    environment {
        DOCKER_CREDS = credentials('docker-hub-credentials')
        IMAGE_NAME = "aadithyaar/cia2-app"
        IMAGE_TAG = "${env.BUILD_ID}"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
        
        stage('Build Docker Image') {
            steps {
                sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG} -t ${IMAGE_NAME}:latest ."
            }
        }
        
        stage('Push to Docker Hub') {
            steps {
                sh "echo \$DOCKER_CREDS_PSW | docker login -u \$DOCKER_CREDS_USR --password-stdin"
                sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}"
                sh "docker push ${IMAGE_NAME}:latest"
            }
        }
        
        stage('Deploy to Azure AKS') {
            steps {
                withKubeConfig([credentialsId: 'azure-aks-kubeconfig']) {
                    sh 'export LC_ALL=C.UTF-8 && export LANG=C.UTF-8 && ansible-playbook -i inventory.ini deploy.yml'
                }
            }
        }
    }
}