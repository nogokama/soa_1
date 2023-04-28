set -ex

bash scripts/stop.sh
bash scripts/build.sh
bash scripts/run.sh
bash scripts/test.sh
bash scripts/stop.sh