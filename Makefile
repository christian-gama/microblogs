.PHONY: post
post:
	go run ./posts/*.go

.PHONY: comment
comment:
	go run ./comments/*.go

.PHONY: event
event:
	go run ./eventbus/*.go

.PHONY: moderation
moderation:
	go run ./moderation/*.go

.PHONY: query
query:
	go run ./query/*.go