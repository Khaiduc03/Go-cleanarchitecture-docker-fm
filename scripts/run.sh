

TYPE=$1
# change_mode() {
#     sed -i "s/MODE=.*/MODE=$TYPE/" .env
#     sed -i "s/TYPE=.*/TYPE=dev/" .env
# }

check_image_exist() {
    local image_name="mujag/fm-fiber:$TYPE"
    if [ $(sudo docker images -q $image_name) ]; then
       echo "Image already exist"
    else
     sudo docker build -t mujag/fm-fiber:$TYPE .
    fi
}

run_prod() {
       yes | cp -rf docker/Dockerfile.prod ./Dockerfile
       yes | CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go
        check_image_exist
}


run_dev() {
     yes | cp -rf docker/Dockerfile.dev ./Dockerfile
    check_image_exist
}

run() {
   sudo docker compose up
}

setup() {
    echo "Setting up $TYPE environment"

    if [ "$TYPE" = "dev" ]; then
        run_dev
        change_mode
        run
    elif [ "$TYPE" = "prod" ]; then
        run_prod
        change_mode
        run
    else
        echo "Invalid mode"
    fi

    echo "Setup complete"
}

setup



