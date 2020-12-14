# MQTT-lab
## Overview 
  This lab creates two **Go** clients that communicate through the **MQTT protocol**.
## Topics

| Publisher     | Topic         | Frequency | QoS|
|:-------------:|:-------------:|:-----:|:--:|
| SM            | consumption   | 1/60 sec |2|
| SM            | production    | 1/15 sec |2|
| MDMS          | price         | 1/60 sec |0|
| MDMS          | production    | 1/1 sec |2|
## Test
You can test this lab using docker-compose.
1. Clone this repository <br>
`$ git clone https://github.com/Nezz7/docker-swarm-lab.git`
2. Change the current working directory <br>
`$ cd /MQTT-lab`<br>
3. Create the docker image : sm-mqtt  <br>
`$ docker build -t sm-mqtt /sm`<br>
4. Create the docker image : mdms-mqtt  <br>
`$ docker build -t mdms-mqtt /mdms`<br>
5. Run the lab <br>
`$ docker-compose up -d`<br>


## Screenshots

![screenshot](https://github.com/Nezz7/MQTT-lab/blob/main/screenshot/screenshot.PNG)
