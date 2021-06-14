FROM python:3.8-slim as serve
WORKDIR /usr/src/app
# Copy all packages instead of rerunning pip install
COPY requirements.txt ./requirements.txt
RUN     pip install -r requirements.txt 
COPY ./flask-api/*.py ./
CMD ["python", "run_app.py"]

