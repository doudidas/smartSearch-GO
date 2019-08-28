import requests
import urllib3
import json
import sys

# Variables
key = "WXGJ-KSHX-FHJH-O2GR"

# Get user arguments

userInput = raw_input("How much value do you want to add ? [default:100]")

if userInput is "":
    print("Default value used: 100")
    userInput = "100"

# Set URL
url = "https://randomuser.me/api?results={}&key={}&fmt=pretty&noinfo".format(
    userInput, key)

# Get databse
try:
    with open("random-database.json", "r") as file:
        database = json.load(file)
except ValueError as e:
    database = []


# Set header
headers = {"Accept": "application/json",
           "Content-Type": "application/json", 'Accept-Encoding': "gzip, deflate"}

# Set Query

querystring = {"results": userInput,
               "key": "WXGJ-KSHX-FHJH-O2GR", "fmt": "pretty", "noinfo": ""}

print(url)
response = requests.request("GET", url, headers=headers)
j = response.json()


def customMap(e):
    return {
        "firstname": e["name"]["first"],
        "lastname" : e["name"]["last"],
        "username" : e["login"]["username"].encode('utf-8'),
        "email"    : e["email"],
        "picture"  : e["picture"]
    }


database = map(customMap, j["results"])

# Send request

with open("random-database.json", 'w') as f:
    json.dump(database, f)
