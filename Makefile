run: consumer/consumer consumer/provider.so
	consumer/consumer -plugin consumer/provider.so

consumer/consumer: consumer/main.go shared/ifi.go
	@cd consumer && \
		go build -o consumer .

consumer/provider.so: provider/main.go shared/ifi.go
	@cd provider && \
		go build -buildmode=plugin -o ../consumer/provider.so .
