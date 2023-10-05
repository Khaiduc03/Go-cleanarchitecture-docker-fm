#/bin/zsh

MODE=$1

setup_dev() {
    yes | cp -rf configs/dev/.env .env
    yes | cp -rf configs/dev/service-account.json service-account.json
}

setup_prod() {
    yes | cp -rf configs/prod/.env .env
    yes | cp -rf configs/prod/service-account.json service-account.json
}

setup() {
    echo "Setting up $MODE environment"

    if [ "$MODE" = "dev" ]; then
        setup_dev
    elif [ "$MODE" = "prod" ]; then
        setup_prod
    else
        echo "Invalid mode"
    fi

    echo "Setup complete"
}

setup