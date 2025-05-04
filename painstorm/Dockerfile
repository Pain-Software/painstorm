FROM postgres:17.4

RUN apt-get update && apt-get -y install git build-essential postgresql-server-dev-17

RUN git clone https://github.com/citusdata/pg_cron.git
RUN cd pg_cron && make && make install  