default:
	docker-compose up --force-recreate --build

build:
	docker-compose build --no-cache

dev:
	docker-compose -f docker-compose-dev.yml up --force-recreate --build

prettier:
	prettier --write . !ui/browser
	git add .
	git commit -m "prettierize"
	git push

prettier-check:
	prettier --check .

release:
	docker-compose build
	docker save iris_ui:latest > iris_ui.tar
	gpg --sign iris_ui.tar
