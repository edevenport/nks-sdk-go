pipeline {
  agent {
    label "jenkins-go"
  }

  triggers {
    cron('H 0 * * *')
  }

  environment {
    ORG            = 'edevenport'
    APP_NAME       = 'nks-sdk-go'
    GO111MODULE    = 'on'
    NKS_API_TOKEN  = credentials('nks_api_token')
    NKS_API_URL    = 'https://api-staging.stackpoint.io/'
    NKS_SSH_KEYSET = 3750
    NKS_AWS_KEYSET = 3751
    NKS_EKS_KEYSET = 5878
    NKS_AZR_KEYSET = 4954
    NKS_AKS_KEYSET = 6511
    NKS_GCE_KEYSET = 6575
    NKS_GKE_KEYSET = 6578
  }

  stages {
    stage('CI Build and push snapshot') {
      when {
        // branch 'PR-*'
        branch 'task/ci_integration'
      }
      steps {
        container('go') {
          dir('/home/jenkins/go/src/github.com/edevenport/nks-sdk-go/nks') {
            checkout scm
            sh "go version"
            sh "go mod vendor"
	    // sh "go get github.com/golang/lint/golint"
	    sh "golint"
	    // sh "go vet"
            sh "go test -v -timeout=120m -run=TestLiveBasicClient"
          }
        }
      }
    }
  }
}
