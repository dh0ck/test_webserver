pipeline {
    agent {
        node {
            label ''
        }
    }

	triggers {
		pollSCM '*/1 * * * *'
	}

    stages {
        stage('Build') {
            steps {
                echo "building"
                sh '''
                echo "doint building sturf"
                '''
            }
        }
    
        stage('Test') {
            steps {
                echo "testing"
                sh '''
                echo "doint test stuff"
                '''
            }
        }
        stage('Deliver') {
            steps {
                echo "delivering"
                sh '''
                echo "doint deliver stuff"
                '''
            }
        }
    }
}
