default:
	docker-compose up --force-recreate --build

build:
	docker-compose build --no-cache

dev:
	docker-compose -f docker-compose-dev.yml up --force-recreate --build

prettier:
	prettier --write .
	git add .
	git commit -m "prettierize"
	git push

prettier-check:
	prettier --check .

release:
	rm -f iris-ui.tar
	docker build -t iris-ui:latest ui
	docker save iris-ui:latest > iris-ui.tar
	gpg --sign iris-ui.tar