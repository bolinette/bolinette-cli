FROM python:3-alpine

ENV WORKER=3

ADD requirements.txt ./

RUN apk add --no-cache py-pip

RUN pip install gunicorn && \
    pip install -r requirements.txt

WORKDIR /opt/{{.Module}}

COPY ./ ./

RUN echo "production" > ./env/.profile

EXPOSE 5000

CMD ["gunicorn", "-w", ${WORKER}, "server:app", "--bind", ":5000", "--worker-class", "aiohttp.GunicornWebWorker"]