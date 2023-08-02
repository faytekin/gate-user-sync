project_name = gate_user_sync

up:
	docker compose -p $(project_name) up -d

down:
	docker compose -p $(project_name) down

logs:
	docker compose -p $(project_name) logs -f gate_user_sync

build:
	docker compose build
