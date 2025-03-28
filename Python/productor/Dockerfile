FROM python:3

WORKDIR /usr/src/app

#COPY requirements.txt ./
RUN pip install pika

COPY . .

CMD [ "python", "./send.py" ]
