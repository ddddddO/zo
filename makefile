build:
	docker build -f dockerfile/ui/Dockerfile -t gcr.io/newmap-1551402398637/ui .
	gcloud config set project newmap-1551402398637
	gcloud auth configure-docker
	docker push gcr.io/newmap-1551402398637/ui

# newmap-1551402398637 にテスト用testddddddoバケットがある
# ref: https://cloud.google.com/run/docs/quickstarts/build-and-deploy/go?hl=ja#deploy

# ref: https://www.apps-gcp.com/cloud-run-full-managed-basic-1/
# 以下、初回gcloud run deploy --image=gcr.io/newmap-1551402398637/ui　実行時のインタラクティブ
# Service name (ui):  zo-ui
# API [run.googleapis.com] not enabled on project [684654458149]. Would
# you like to enable and retry (this will take a few minutes)? (y/N)?  y
# ...
# Please specify a region:
# [1] asia-east1
# ↓
# 1
# ...
# FIXME: 以下が問題。一覧画面までの認証をどうするか。
# Allow unauthenticated invocations to [zo-ui] (y/N)?  N  <- で、ブラウザから行くと「Error: Forbidden」。認証を求められることもなかった。ので、一旦、オープンにしてアクセスした。
# ...
# Service URL: https://zo-ui-kqfxvbdlza-de.a.run.app
deploy:
	gcloud config set project newmap-1551402398637
	gcloud run deploy --image=gcr.io/newmap-1551402398637/ui
