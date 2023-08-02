project_name = gate_user_sync

up:
	docker compose -p $(project_name) up -d

down:
	docker compose -p $(project_name) down

build:
	docker compose build
