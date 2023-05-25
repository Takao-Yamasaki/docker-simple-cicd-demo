up-prod: # Do docker compose up for production
		docker compose -f docker-compose-prod.yml up

up-prod-build: # Do docker compose up for production with rebuild
		docker compose -f docker-compose-prod.yml up --build
