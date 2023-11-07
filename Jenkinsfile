pipeline {
    agent any
        triggers {
                githubPush()
            }
    options {
            buildDiscarder logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '10', daysToKeepStr: '', numToKeepStr: '10')
        }
    stages {
        stage('Checkout') {
            steps {
                git credentialsId: 'amanpd-github-credentials', url: 'https://github.com/authnull0/user-service.git', branch: 'signup-api'
            }
        }
        stage('Sonarqube Scanning') {
            environment {
                    scannerHome = tool 'SonarQubeScanner'
                    scannerCmd = "${scannerHome}/bin/sonar-scanner"
                    scannerCmdOptions = "-Dsonar.projectKey=user-service -Dsonar.sources=app,conf,src,utils -Dsonar.host.url=http://195.201.165.12:9000"
                }
            steps {
                withSonarQubeEnv(installationName: 'sonarqube-server') {
                sh "${scannerCmd} ${scannerCmdOptions}"
                }
                timeout(time: 10, unit: 'MINUTES') {
                waitForQualityGate abortPipeline: true
                }
            }
        }
        stage('Scan for git-secrets') {
            steps {
                sh 'git secrets --scan -r'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    def dockerImage = "user-service:${env.BUILD_ID}"
                    env.dockerImage = dockerImage
                    sh "docker build . --tag ${dockerImage}"
                    echo "Docker Image Name: ${dockerImage}"
                    echo "${env.BUILD_ID}"
                }
            }
        }
        stage('Trivy Scan Docker Image') {
            steps {
                script {
                    def formatOption = "--format template --template \"@/usr/local/share/trivy/templates/html.tpl\""
                    sh """
                    trivy image ${env.dockerImage} $formatOption --timeout 10m --output report.html || true
                    """
            }
        publishHTML(target: [
          allowMissing: true,
          alwaysLinkToLastBuild: false,
          keepAll: true,
          reportDir: ".",
          reportFiles: "report.html",
          reportName: "Trivy Report",
        ])
            }
        }
        stage('Stop & Remove older image') {
            steps {
                script {
                    def currentBuildNumber = env.BUILD_ID.toInteger()
                    
                    // Loop through previous builds from n-1 down to 1 till finds any running container
                    for (int buildNumber = currentBuildNumber - 1; buildNumber >= 1; buildNumber--) {
                        def dockerCommand = "user-service-${buildNumber}"
                        def buildStatus = sh(script: "docker ps -a | grep ${dockerCommand}", returnStatus: true)
                        
                    if (buildStatus == 0) {
                            sh "docker stop ${dockerCommand}"
                            sh "docker rm ${dockerCommand}"
                            break
                        } 
                    }
                }
            }
        }
        stage('Run Docker Container') {
            steps {
                script {
                    def dockerCommand = "-itd --network host --name=user-service-${env.BUILD_ID} ${env.dockerImage}"
                    sh "docker run ${dockerCommand}"
                }
            }
        }
    }
}