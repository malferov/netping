#!/bin/bash
export BOT_TOKEN=$(cat ../../.key/slack.tok)

./send -stderrthreshold=INFO

exit

curl -XPOST https://slack.com/api/chat.postMessage \
  --data '{"channel":"general", "text":"Subject: sub\r\nMessage: msg\r\nEmail: e-mail\r\n"}' \
  --header "Content-Type: application/json; charset=utf-8" \
  --header "Authorization: Bearer $BOT_TOKEN" \
  -v | jq
