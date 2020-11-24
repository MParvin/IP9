FROM python:rc-alpine3.12

RUN pip install Flask

WORKDIR /app

COPY . .

EXPOSE 5000

ENTRYPOINT ["python","flask_app.py"]
