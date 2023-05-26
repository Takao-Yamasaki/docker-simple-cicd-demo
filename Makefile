build-go:
		docker image build --target production -t temp_api_image_name_prod:latest ./api/

run-go:
		docker container run -it temp_api_image_name_prod:latest sh

up-prod: # Do docker compose up for production
		docker compose -f docker-compose-prod.yml up

up-prod-build: # Do docker compose up for production with rebuild
		docker compose -f docker-compose-prod.yml up --build
