pipeline {
  agent {
    label "jenkins-go"
  }
  environment {
    ORG = 'edevenport'
    APP_NAME = 'nks-sdk-go'
  }
  stages {
    stage('CI Build and push snapshot') {
      when {
        // branch 'PR-*'
        branch 'task/ci_integration'
      }
      steps {
        container('go') {
          dir('/home/jenkins/go/src/github.com/edevenport/nks-sdk-go') {
            checkout scm
            sh "go version"
            // sh "make linux"
            // sh "export VERSION=$PREVIEW_VERSION && skaffold build -f skaffold.yaml"
            // sh "jx step post build --image $DOCKER_REGISTRY/$ORG/$APP_NAME:$PREVIEW_VERSION"
          }
        }
      }
    }
  }
}
