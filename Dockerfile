FROM rabbitmq
MAINTAINER Nathan Craddock <nkcraddock@gmail.com>
ADD ./bin/ /smsgs/bin/
ADD ./scripts/ /smsgs/bin/

# Define environment variables.
ENV RABBITMQ_LOG_BASE /data/log
ENV RABBITMQ_MNESIA_BASE /data/mnesia

# Define mount points.
VOLUME ["/data/log", "/data/mnesia"]

# Define working directory.
WORKDIR /data

CMD ["/smsgs/bin/start.sh"]

EXPOSE 80
EXPOSE 15672

