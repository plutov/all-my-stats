#!/bin/bash

CONFIGURATION_NAME="packagemain"
PROJECT_NAME="alexsandbox"
REGION="europe-west1"

YT_API_KEY=${YT_API_KEY:=""}
YT_CHANNEL_ID="UCI39wKG8GQnuzFPN5SM55qw"
GH_USER="plutov"
TWITTER_USER="pliutau"
GA_ID="115234933"
SOF_ID="350294"

gcloud config configurations activate $CONFIGURATION_NAME

gcloud config set project $PROJECT_NAME
gcloud config set functions/region $REGION

gcloud services enable cloudfunctions.googleapis.com

export GO111MODULE=on
go mod tidy
go mod vendor

gcloud functions deploy stats --entry-point Stats \
--runtime go111 \
--trigger-http \
--set-env-vars YT_API_KEY=$YT_API_KEY \
--set-env-vars YT_CHANNEL_ID=$YT_CHANNEL_ID \
--set-env-vars GA_ID=$GA_ID \
--set-env-vars GH_USER=$GH_USER \
--set-env-vars TWITTER_USER=$TWITTER_USER \
--set-env-vars SOF_ID=$SOF_ID