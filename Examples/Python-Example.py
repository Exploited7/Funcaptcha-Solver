import requests

def SolveCaptcha(publickey,website,host,blob):
  data = {
    "host": host,
    "publickey": publickey,
    "website": website,
    "blob": blob,
    "solvekey":"TM-upBsYRUYH4SfbpQ13j0AuBpf7giICLVBb5yDDPZYwDSjs"
  }
  Solve = requests.post('http://23.137.104.216:5000/api/funcaptcha',json=data)
  print(Solve.json()['token'])

SolveCaptcha('73BEC076-3E53-30F5-B1EB-84F494D43DBA',"https://signin.ea.com",'ea-api.arkoselabs.com',"undefined)
