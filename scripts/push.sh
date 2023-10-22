

check_image_exist() {
    local image_name="mujag/fm-fiber:prod"
    if [ $(sudo docker images -q $image_name) ]; then
       echo "Image already exist"
    else
     sudo docker build -t mujag/fm-fiber:prod .
     sudo docker push mujag/fm-fiber:prod
    fi
}


check_image_exist

