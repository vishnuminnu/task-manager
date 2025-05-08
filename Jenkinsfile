pipeline {
    agent any
    
    tools {
        nodejs 'Node16'  // Make sure this is configured in Jenkins Global Tools
        go 'Go'          // Make sure this is configured in Jenkins Global Tools
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
                    bat 'npm install'  // Changed from sh to bat for Windows
                    bat 'npm run build'
                }
            }
        }
        
        stage('Build Backend') {
            steps {
                dir('backend') {
                    bat 'go build -o taskmanager.exe'  // Changed from sh to bat
                }
            }
        }
    }
}