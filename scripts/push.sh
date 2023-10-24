
change_mode() {
    sed -i "s/TYPE=.*/TYPE=prod/" .env
}


check_image_exist() {
    change_mode
    if [ $(sudo docker images -q mujag/fm-fiber:prod) ]; then
       echo "Image already exist"
        sudo docker push mujag/fm-fiber:prod
    else
     sudo docker build -t mujag/fm-fiber:prod -f Dockerfile.prod .
     sudo docker push mujag/fm-fiber:prod
    fi
}


check_image_exist

