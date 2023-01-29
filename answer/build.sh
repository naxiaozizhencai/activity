/usr/local/go/bin/go build -o answer_api  ./api/answer.go
nohup ./answer_api -f /data/app/activity/answer/api/etc/answer-api.yaml >  answer_api.out &