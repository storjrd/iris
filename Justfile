default:
	DOCKER_BUILDKIT=1 docker-compose up --force-recreate --build

build:
	DOCKER_BUILDKIT=1 docker-compose build --no-cache

dev:
	DOCKER_BUILDKIT=1 docker-compose -f docker-compose-dev.yml up --force-recreate --build

test:
	mkdir -p ui-tests-output
	DOCKER_BUILDKIT=1 docker-compose -f docker-compose-test.yml build --parallel
	DOCKER_BUILDKIT=1 docker-compose -f docker-compose-test.yml up --force-recreate --abort-on-container-exit

prettier:
	prettier --write . !ui/browser
	git add .
	git commit -m "prettierize"
	git push

prettier-check:
	prettier --check .

update-browser:
	git pull
	cd ui/browser && git fetch; git checkout main; git pull
	git add .
	git commit -m "bump browser version"
	git push
