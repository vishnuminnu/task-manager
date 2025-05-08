pipeline {
    agent any
    
    // PROPER tools section (no trailing characters)
    tools {
        nodejs 'Node16'  // Must match exact name in Jenkins Global Tools
        go 'Go'          // Must match exact name in Jenkins Global Tools 
    }
    
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', 
                url: 'https://github.com/vishnuminnu/task-manager.git'
            }
        }
        
        stage('Build Frontend') {
            steps {
                dir('frontend') {
                    sh 'npm install'
                    sh 'npm run build'
                }
            }
        }
        
        stage('Build Backend') {
            steps {
                dir('backend') {
                    sh 'go build -o taskmanager'
                }
            }
        }
    }
}