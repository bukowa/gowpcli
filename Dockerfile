FROM wordpress:cli-php7.4
WORKDIR /app
USER root
ENTRYPOINT ["/bin/bash", "-c"]
CMD ["wp --allow-root cli cmd-dump > dump.json"]
