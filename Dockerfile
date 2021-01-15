FROM wordpress:cli-php7.4
WORKDIR /app
ENTRYPOINT []
CMD ["/bin/bash", "-c", "wp cli cmd-dump > dump.json"]
