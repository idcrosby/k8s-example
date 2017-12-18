node {
    checkout scm
    def DOCKER_HUB_ACCOUNT = 'icrosby'
    def DOCKER_IMAGE_NAME = 'k8s-example-adidas'

    echo 'Building Go App'
    stage("build") {
        docker.image("icrosby/jenkins-agent:kube").inside('-u root') {
            sh 'go build' 
        }
    }
    echo 'Testing Go App'
    stage("test") {
        docker.image('icrosby/jenkins-agent:kube').inside('-u root') {
            sh 'go test' 
        }
    }

    echo 'Building Docker image'
    stage('BuildImage') 
    def app = docker.build("${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}", '.')

    echo 'Testing Docker image'  
    stage("test image") {
        docker.image("${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}").inside {
            writeFile file: '/test.sh', text: "#!/bin/bash
set -eu

/home/server &
ID=$! # ID of webserver process, so we can kill it

tests_passed=true
expected="Hello From Adidas."
output=$(curl -s localhost:8080)
if [[ $output == *"$expected"* ]]; then
  echo "Test Success"
else
  echo "Test Failure"
  echo "$expected != $output"
  tests_passed=false
fi


kill $ID

if [[ "$tests_passed" == "true" ]]; then
  echo "Passed Tests"
else 
  echo "Failed Tests"
  exit 1
fi"
            sh './test.sh'
        }
    }
}
