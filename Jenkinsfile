node {
    stage('Get Code'){
git url: 'https://github.com/byebye758/web-load-test.git'
}


stage('Build package to  Docker'){
   def newApp = docker.build "registry.cloud.ixicar.cn/library/webhook:${BUILD_TAG}"
   newApp.push()
    
}


stage('Push  kubernetes'){
sh 'autobox --kubectlpath=/usr/local/bin/kubectl  --Projectname=${JOB_NAME} --kubeconfigpath=/etc/kubernetes/admin.conf --replace=1  --image=registry.cloud.ixicar.cn/library/webhook:${BUILD_TAG}  --port="8082:8082"  --http="url=webwook.cloud.ixicar.cn,serviceport=8082,https=false,path=/" --autoscal="min=1,max=20,cpuload=50m"  '

}
}