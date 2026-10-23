pipeline {
    agent any
    
    environment {
        // These credentials must be configured in Jenkins
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
                // Requires Kubernetes CLI Plugin and kubeconfig credentials in Jenkins
                withKubeConfig([credentialsId: 'azure-aks-kubeconfig']) {
                    sh "kubectl apply -f k8s.yaml"
                    sh "kubectl rollout restart deployment/cia2-app"
                }
            }
        }
    }
}