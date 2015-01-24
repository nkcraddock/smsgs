FROM rabbitmq:3.4.3

RUN rabbitmq-plugins enable --offline rabbitmq_management
RUN rabbitmqctl add_user smsgs parseword
RUN rabbitmqctl set_permissions smsgs ".*" ".*" ".*"


EXPOSE 15672
