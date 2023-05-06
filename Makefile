
.PHONY: deploy
deploy:
	gcloud functions deploy SmsToSlackHandler --runtime go120 --trigger-http --allow-unauthenticated --entry-point SmsToSlackHandler