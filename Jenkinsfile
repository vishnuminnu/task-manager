pipeline {
    agent any
    tools {
        nodejs 'Node16' // Node.js installation defined in Global Tool Configuration
        go 'Go1.21'     // Go installation defined in Global Tool Configuration
    }
    stages {
        stage('Checkout') {
            steps {
                // Clone the GitHub repository
                git url: 'https://github.com/your-username/task-manager.git', branch: 'main'
            }
        }
        stage('Build Frontend') {
            steps {
                // Navigate to frontend directory and build React app
                dir('frontend') {
                    sh 'npm install'
                    sh 'npm run build'
                }
            }
        }
        stage('Build Backend') {
            steps {
                // Navigate to backend directory and build Go application
                dir('backend') {
                    sh 'go mod tidy'
                    sh 'go build -o task-manager-backend'
                }
            }
        }
        stage('Archive Artifacts') {
            steps {
                // Archive frontend build and backend binary
                archiveArtifacts artifacts: 'frontend/build/**, backend/task-manager-backend', fingerprint: true
            }
        }
    }
    post {
        success {
            echo 'Build completed successfully!'
        }
        failure {
            echo 'Build failed. Check the logs.'
        }
    }
}