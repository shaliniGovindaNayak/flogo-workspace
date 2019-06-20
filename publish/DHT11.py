#!/usr/bin/python
import Adafruit_DHT
import json
import paho.mqtt.client as mqtt


client = mqtt.Client()
client.username_pw_set("username","password")
client.connect("192.168.1.59",1883,60)
sensor = Adafruit_DHT.DHT11
pin = '17'

while True:
  if humidity is not None and temperature is not None:
   
        humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
        data={}
        data["temperature"]=temperature
        data["humidity"]=humidity
        json_Data=json.dumps(data)
        print(json_Data)
        client.publish("test", json_Data)
        print("published")
        client.disconnect()


  else:
        print('Failed to get reading. Try again!')

