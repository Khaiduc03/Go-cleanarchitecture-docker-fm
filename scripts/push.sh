
change_mode() {
     sed -i "s/TYPE=.*/TYPE=prod/" .env
}


check_image_exist() {
    sudo docker rmi mujag/fm-fiber:prod -f
     sudo docker build -t mujag/fm-fiber:prod -f Dockerfile.prod .
     sudo docker push mujag/fm-fiber:prod
}
change_mode

check_image_exist

