import time
import requests
import os

def download(localFile, srcUrl):
	print("------------------------------------------------------------")
	print('Downloading %s' % localFile, end='\r')
	with requests.get(srcUrl, stream=True) as r:
		if r.status_code != 200:
			print("retrying")
			return download(localFile, srcUrl)
		contentLength = int(r.headers['content-length'])
		print('Downloading %s %.2f MB' % (localFile, contentLength/1024/1024))
		downSize = 0
		startTime = time.time()
		with open(localFile+".temp", 'wb') as f:
			for chunk in r.iter_content(8192):
				if chunk:
					f.write(chunk)
				downSize += len(chunk)
				line = '%.1f%% - %.2f MB/s - %.2f MB          '
				line = line % (downSize/contentLength*100, downSize/1024/1024/(time.time()-startTime), downSize/1024/1024)
				print(line, end='\r')
				if downSize >= contentLength:
					break
		os.rename(localFile+".temp", localFile)
		print()

Block_Files = {
	"./pkg/data/0to999999_Block.zip":			"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVExUUJSVnpUQk5IdmZSNWxxZTBlVnNCSzNTSXlLakk5WXBFSlpvN2RXNFkwUT9lPVg0UmFXTw==.zip",
	"./pkg/data/1000000to1999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVNNSmpwUlJzalJOcFdVamx5OWVWLTRCcDVfUWNIejNLcjVwN2djYjB3dDY2Zz9lPVVzdkdudQ==.zip",
	"./pkg/data/2000000to2999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVRCMEkzNlhEeUZEbEJtd255blRlQVlCWWxzZUFid2pDNl9yb1BWSWRZWm5NUT9lPTdiUmgzTw==.zip",
	"./pkg/data/3000000to3999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVhFLU5ydGZwVzlGdFNiRExxYmF5M2NCaDl2RkNwMjhyeE9EWjVxMW9MQWdCQT9lPUJxcG80Yg==.zip",
	"./pkg/data/4000000to4999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVl5ZUluY1FVaDlBcVIyb01YWlU0VU1CVXA1Ty1QOVVZOHY3YVgtTmVxdk1zQT9lPTVvZmQxcg==.zip",
	"./pkg/data/5000000to5999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVYwbWJoM09JdmhIdG82Tk9SdzRkMThCR0xzYVRrSjRsR0hXRlV5Zjc4Ml9TQT9lPVdHZGhvbQ==.zip",
	"./pkg/data/6000000to6999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVNMcWZsLUJnMzlJaWE4WFh1VmVhSVVCVjdIQ1loa3RMVWlkZ3BYU19ULXZrZz9lPTJUUlV0Sw==.zip",
	"./pkg/data/7000000to7999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVJyVjRMSlRxak5CdkkyRmhIMU9EdWdCTnFlU204VDhCbHdqdHdFbGZxTU9Ldz9lPUdLSk1vcA==.zip",
	"./pkg/data/8000000to8999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRzcE96dTFBZnBLcm5EVGpTU01MazBCRVNIdEJyU3U5V214NHdIRWhlOC1BUT9lPVNoelllbw==.zip",
	"./pkg/data/9000000to9999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWJKOUtaejdpVmhMaW0wdUwzTHpOTEFCbVdKREVyQWZKSndiQ0kxdDFXR3J5Zz9lPVJIRmdEaA==.zip",
	"./pkg/data/10000000to10999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVg4dHpaLWNlM3hPdnJHNmdxRXFXU3dCM3FXLVpSQ0VVS0xmWTlvYjdkNXA4QT9lPWpzQmhiSw==.zip",
	"./pkg/data/11000000to11999999_Block.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVpWcFFtMDkxYlJQdmp6alVLS0h2S2dCeGdjSHV0TG1HUWRSTVN4SlQ3MFFTdz9lPUMyZ1VRdQ==.zip",
	"./pkg/data/12000000to12999999_Block.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVNkSWY3VGc3TDlLcG50aDNidlZjd3NCbXNJRVNHbjU4SFBnOGZPMzRFNzFhZz9lPUNmV1FqcQ.zip",
	"./pkg/data/13000000to13249999_Block.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRKb1pqbUdScDlCdkdTT1VSMWI1Z2NCVXI0eEJPanBPeVZZX3FvRFUzWE96dz9lPU9keGJHMQ.zip",
	"./pkg/data/13250000to13499999_Block.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVJpWENYSmFWZE5CaFl1NlBIYW9kWGtCb0VBTEpRZG50cEtzZlNRZGt5Wk5KUQ.zip",
	"./pkg/data/13500000to13749999_Block.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWM1b2JhY3JUaUJJdTAwTzQ2Vk5zZ2NCY2piSG1waHZ5N2RNcUVCWUtpUEViUT9lPTNBb1dHMg.zip",
	"./pkg/data/13750000to13999999_Block.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVdlS3J5Z0VfQzVNbnJ5YlpQY0dIQVlCQnN1NS16NTZQMlRPMGl4ZDc4QVBVZz9lPXJnZmpQWg.zip",
	"./pkg/data/14000000to14249999_Block.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVZPR0ZpWFpDWjVKb3FOVDU0NTA4UXdCTFU2dEtpQmNVcXdCYVVsdC1MemNPUT9lPU1wd3V1dw.zip",
	"./pkg/data/14250000to14499999_Block.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWJyUnlaTXZqb3hGbnhOVFIyV3hncWdCSzQyck11aFhIQnZCRHMzaHEyaGhlQT9lPURUSFVDQQ.zip",
	"./pkg/data/stat_Block.py":				"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVluR1UzY0ZLakJHdnVwMUxYcFhtUVlCQ0UtZG1fRDFOcFdENGxqenVQM2NEdz9lPXdqTEtibA==.py"
}
BlockTransaction_Files = {
	"./pkg/data/0to999999_BlockTransaction.zip":			"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWNMZXJ0bU4xazFMcG51WVZSTmRHTllCbVk2YjB0OUFRTjJ4Yk9SUllkekVOZz9lPVRVdTh6cQ==.zip",
	"./pkg/data/1000000to1999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVp3Rk9XTERHMlJCci1Xa293NC1IdndCdHNEWjd4VXhpeFhMYXQ1OVZoUzgyQT9lPU16dnZKag==.zip",
	"./pkg/data/2000000to2999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVJ5b2VadjF4bWhCcnNxZzRfdjZVc29CRnhISGNERzJfWXZweXo3bEl1NVZRdz9lPXlpYllaVw==.zip",
	"./pkg/data/3000000to3999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWM5RjI0R1M5NEpHaW9LbW83WjZmajRCZTFqMGJQZVUydFBUQWVKU0FzbXRzUT9lPU4zYnZ0cw==.zip",
	"./pkg/data/4000000to4999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVpucXh2bF9sRmhGa012X2NaZnI3UklCQ1V3S0RDdXBnZE56bFVXYmVqVFJrZz9lPTBCTXJuQQ==.zip",
	"./pkg/data/5000000to5999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVZHZEpuYmJpRjFDdHlTbXFfcUxtU1lCdlRDSGNTY0xmV2hyazBYbl8yeE1udz9lPTZlQ3M5TA==.zip",
	"./pkg/data/6000000to6999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRITk9jTEVkZmhBbDE0WlV2Y0VfaUFCVmFBT2M4c2dKdnJyV0VTaWw2eGhKQT9lPWhSSXJ4SQ==.zip",
	"./pkg/data/7000000to7999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWFMOTE0ZWhFSE5CcnZjcEd5WUZTQW9Cd3JJYmMtS3BZd05JX3h1b3pHVUFnQT9lPUVFZEc5cA==.zip",
	"./pkg/data/8000000to8999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVQxcVpCMHdIOWhNZ2xiWmkycDl0Q1lCNmNBaWQ3QlhoVklJMGhTakxXYjR5dz9lPVIzaEZjVQ==.zip",
	"./pkg/data/9000000to9999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVdiS3RZcl9acEZFdGdPMm9JWE1kX1FCMFVOVUJ2bnBIODVneWV2VU5Tajhydz9lPXZiQmZNRw==.zip",
	"./pkg/data/10000000to10999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWY5aWlLOHFaNmRQazNvSzQzM1pJSElCV1JENlJ6NFJpbzhGNmlhd2xUeUFXQT9lPVNvUng5SQ==.zip",
	"./pkg/data/11000000to11999999_BlockTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVRGQkJ0X3BPYkZDbWd6U005eHlYNjRCWkNlUVNBeVplMGdKeTRkQWhxRnVIZz9lPXRza1JPRg==.zip",
	"./pkg/data/12000000to12999999_BlockTransaction.zip": 	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVJrM1l1VWZoWnBMaF9OUjdoZnZydHNCalA4M1dmU2gtSEEtdEFiNE9fVXZ6dz9lPWJnWmticw.zip",
	"./pkg/data/13000000to13249999_BlockTransaction.zip": 	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWNfU28xeEUtMTVCdkduanpoLXg5UTRCS2JCQlZFMHZsSXpWWTNvNEprbS1Ldz9lPWFra1lNcw.zip",
	"./pkg/data/13250000to13499999_BlockTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRPTHRnYzhWXzFMdm1Fb1dRWkFsLXNCZWJlY0RSVERBaTVGNm8teEV3MUpXQT9lPWFQTlRycw.zip",
	"./pkg/data/13500000to13749999_BlockTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVphcE55Q3hhTk5Hb0hFMDkyUnFQcVlCMkRyYTc5c2NIQ3NQNnJZNGxHWWtOQT9lPXVpaGliQg.zip",
	"./pkg/data/13750000to13999999_BlockTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWVjcUVJYUdpRWhGcHNIdzFtYjY3ckVCeS1fLS1XQlR4bS1BSG4weEprU2t4UT9lPXNaaTVaYQ.zip",
	"./pkg/data/14000000to14249999_BlockTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWZjdzItV1BITXBKaHQ2R3EzSTN2ZGtCeU9UYWxjaGk0X0ZjTVQzYzNpUkRXdz9lPWppanpGUQ.zip",
	"./pkg/data/14250000to14499999_BlockTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRma0FyVXVvWU5KcWNFdVdWczhOOEVCSkYwaEhKMGVramVRSkJTOXVvY3p4Zz9lPWRYZGlKQg.zip",
	"./pkg/data/stat_BlockTransaction.py":					"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRLRjZuSDgzQnRIcWJQXzE5U0psN1lCMEJvYV9VWElNR0plcGdwXy1pWk9WZz9lPXF2eGZzOQ==.py"
}
InternalTransaction_Files = {
	"./pkg/data/0to999999_InternalTransaction.zip":			"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVgzOXdVU19fNHhMaUlJLTl5WFEtVjRCcExydG1EdWR2X2FINVhNeUdGcVBFdz9lPUVlV2Y2dw==.zip",
	"./pkg/data/1000000to1999999_InternalTransaction.zip":		"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVNhQU44Y1AyYVZOaWJvUDBYcGpKcUVCbUhlUEFpcjdJcEV0dVk2TzlsU1NiZz9lPUw5ZXdiaw==.zip",
	"./pkg/data/2000000to2999999_InternalTransaction.zip":		"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWFLYlRELWlDYXROcmZJUTE3bTZLeU1CQWZ4RFo0UGJWWXJxcUpGRW1WcTU0QT9lPUR1WE5yNw==.zip",
	"./pkg/data/3000000to3999999_InternalTransaction.zip":		"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVRTZG1HQXlDYjFQcXpwQ0hxSTQtcllCX3lIc1hFZGlCNFBiYlJCRnV6ci1ndz9lPWNjRmRVWA==.zip",
	"./pkg/data/4000000to4999999_InternalTransaction.zip":		"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVo1SzlBcjY3bTlDcHBzVlctaC0tRmtCaG4tTEFLd1BDYURGNzdreG14Yk5Wdz9lPWZhOWNaeg==.zip",
	"./pkg/data/5000000to5999999_InternalTransaction.zip":		"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWJlLVU0RW5VYnBIalVvVUdaOE50ZUlCNHJ0OFhPaTlCenAtV3V3ZWJhTy1KZz9lPVcxZTJ0Tw==.zip",
	"./pkg/data/6000000to6999999_InternalTransaction.zip":		"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWR0ZnpFTmRmczFJdENOS08tMTNtQ0VCdnFFNll1NEZpbU9wa0wzbFRudXFaZz9lPWM3YXpjOA==.zip",
	"./pkg/data/7000000to7999999_InternalTransaction.zip":		"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVZkdEliaV9lSDlDblg1RlFNSXRNZHNCbm95RERCR044ZkdYR0NtOWh3MUVxZz9lPWFoU1BLbA==.zip",
	"./pkg/data/8000000to8999999_InternalTransaction.zip":		"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVRrdFZfWTB0NHBOaDlQdkVPRWx3WFFCOGQ4Q2V2ZmZyc1ExRzFTcDlQOHJJQT9lPVdsU0dZTw==.zip",
	"./pkg/data/9000000to9999999_InternalTransaction.zip":		"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWUtMklUSDZMSmxCdFBmTTVwZEZObjBCUWhUelhnUXhERHNiSXV5ODc3aFdVZz9lPTk5MWE1Nw==.zip",
	"./pkg/data/10000000to10999999_InternalTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVJvMUZ6dmF3bWhQazBybGw3ci1EMXdCaDlMUUlOVHN1QlNpempLclpfTjVVQT9lPTBucEVoQw==.zip",
	"./pkg/data/11000000to11999999_InternalTransaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVE3X1R4SGowNlZDckJhWE50OEVIQUFCWndiVWJXcVRNbGM4QmhKUVZVYnJkQQ==.zip",
	"./pkg/data/12000000to12999999_InternalTransaction.zip": 	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVdnN0duT1hYRjFGanRYT0gzTGNWVU1CY1hTNzMxSXdSRFR5eVNwN1BtVUNHQT9lPUZZckRKNw.zip",
	"./pkg/data/13000000to13249999_InternalTransaction.zip": 	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWNpYW1VUlBxbDVLbnFTM2FJblg2R0FCMU5BQUNCbzZKSUhMY1BUcFhKRlpXZz9lPUJ3NzRNTQ.zip",
	"./pkg/data/13250000to13499999_InternalTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWZaN3h3UDV0ZVpQdVFSN3FKOXdvTElCOFR1YWhOZGt5OWFOYlh6QUFmNjBzdz9lPWpvV3Jnbg.zip",
	"./pkg/data/13500000to13749999_InternalTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWVHekZqM0NPTVJKbVNjVUdEQjFtU1lCcFFsVmM1TnNKZndjN3BmQ0kxbFpsdz9lPXNEUWxyZw.zip",
	"./pkg/data/13750000to13999999_InternalTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVhxOGw3YXBRU3hQdXRCcmZCYmRSbGtCZzlobkhwdXBsaUhrakIxWGlxbGxIdz9lPTZMVVNCYw.zip",
	"./pkg/data/14000000to14249999_InternalTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVR5ZnpMNXFxM1ZLdU50V215OFlFRWNCM21LQjlDdFdxY1ZURlNha3ZaaENwUT9lPXlWNVRuYw.zip",
	"./pkg/data/14250000to14499999_InternalTransaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVVob0hLclVrV2xLa1VUY19UUDVSR1FCV3JtMVpFVVlGdmdIekpHWWgxUVMxdz9lPVlKUGV3MQ.zip",
	"./pkg/data/stat_InternalTransaction.py":					"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVVBMnhtTXpFQ0JNa19NZlBJRk5CeGdCRlg5d3I4dGxkRVVINndfS1NOcXNyQT9lPVd5ZVFVUA==.py"
}
ContractInfo_Files = {
    "./pkg/data/0to999999_ContractInfo.zip":          "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWZ6d0tHUVQyOFJGdGNIOUlueTlKSzhCRC0wQ1djWDlvR25MSC1FeldTUEY2Zz9lPW94bGFEZg==.zip",
    "./pkg/data/1000000to1999999_ContractInfo.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWZTVGRybnk5cTVLcW45UVZHVk4yZGdCd085dV9wRU1BcDFiY0pEUGp6a3lKdz9lPXpGSkxrUg==.zip",
    "./pkg/data/2000000to2999999_ContractInfo.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVFlVlMxV3NDbVJJajNLRm5udFpocjhCXzBHWHBpT1I5RU1XRGMxamd5R1NVZz9lPTkyUlBDMw==.zip",
    "./pkg/data/3000000to3999999_ContractInfo.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWY1VFdxSTdMWlJJbllZLTVuQXI3UmtCWWtYdjJrYnNoMHZkY3l4WWhMeEdWQT9lPUpRMWlBVg==.zip",
    "./pkg/data/4000000to4999999_ContractInfo.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWJEX0tRbjVsT0JBb2UxMlhtQmdjOUFCMFU4QmxoSVVONnRjeTBTYWx3TzdBdz9lPUlqNEdzdg==.zip",
    "./pkg/data/5000000to5999999_ContractInfo.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVkzU0daVm9xVXRKdV9vd25OVE5BMnNCQlRfNEk1U2hmdHBlUnR6RGxyenVFZz9lPUVQdmdzVQ==.zip",
    "./pkg/data/6000000to6999999_ContractInfo.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVd5RktzRk5DX2RQcFk3REJhOHMwS3NCYlNlUEdCVkY0MjZqdGxnUGFOeXlaUT9lPVpuc05kQg==.zip",
    "./pkg/data/7000000to7999999_ContractInfo.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRzY3QwZzJ4aWRJak43VjdrbmoydFVCRGxkcWFaS1dfUlVPdFhNNTRQSnZhZz9lPU02TTcwRQ==.zip",
    "./pkg/data/8000000to8999999_ContractInfo.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVVEdDA2RTBidlZJcnByZXpESGlYQTRCUEpGMW5Nd19kdm1JX0ZObk55WVQwZz9lPTVxdmhzbA==.zip",
    "./pkg/data/9000000to9999999_ContractInfo.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVdBOVQzVUVwYjFIaDJYNFNfeW83UjhCVTQ2Nko0UldzWHBvS25FaUc3c1FFdz9lPUhHSzlDOA==.zip",
    "./pkg/data/10000000to10999999_ContractInfo.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWZoRkszUllUN2RNaUpvNzFfUUN3aUVCVHlLTnItQmZpTGtsWElYUnp3c0NGQT9lPTFRcXNEUw==.zip",
    "./pkg/data/11000000to11999999_ContractInfo.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWFTQzNpWVd3WDVMaHNmSzNOd2lRaUlCd2g0Mi1CZGVta0ZVakJuZE5kV29zdz9lPUhTV1owVA==.zip",
	"./pkg/data/12000000to12999999_ContractInfo.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRybjdZVTZsOUJQZ2N3OFlaMVRSaEFCVHhTNFlBV1Q2Sk1zMW0xZHpDc0RWZz9lPUx0dEhWNw.zip",
	"./pkg/data/13000000to13249999_ContractInfo.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVdsNVNlU2I2QVZDcEQzeHFqNGpLT1FCejdic25Mc09PVmtIU1Z1LVNaWUpWdz9lPURDRXRoTQ.zip",
	"./pkg/data/13250000to13499999_ContractInfo.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVVJc0JvVDgzYTVDc21GaTZXWDJ6MzhCMUU4SkVYdi0tOUdrd2llZmNCY1FjUT9lPTJRd2VPUQ.zip",
	"./pkg/data/13500000to13749999_ContractInfo.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVZRNmJVOG95dVZNaHBkZnR6LUM4TThCYnpTY0pXOHAycmlGTUVadXZsYWRqUT9lPTBkSWl0RQ.zip",
	"./pkg/data/13750000to13999999_ContractInfo.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWVlT3lhWkd5REZGdFg3X2FNdUdwZ01COUxyV1dqclFxbnlCN0lxaVV2QzBUZz9lPXg0dVV2Qg.zip",
	"./pkg/data/14000000to14249999_ContractInfo.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVRteV9SQ2hMT3RJdjFnVFZfNElTM3dCaGhScEl0WjhPZ2NETUUtZzJtenkwQT9lPUdMakdlag.zip",
	"./pkg/data/14250000to14499999_ContractInfo.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWY2RGNQSDNTS3RNcXZ1dnFQNmwyTThCQzJqWk54QjJuT2FLdExGNWVEM05QQT9lPVp0NkVFRA.zip",
	"./pkg/data/stat_ContractInfo.py":				   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVk5ZmZ4bWt1QkZOazNiVmNhakQzTVlCSXR2bVA4Sm12dnhCWEtSNXV0S3c3dz9lPXVkb0NzOQ==.py"
}
ERC20Transaction_Files = {
    "./pkg/data/0to999999_ERC20Transaction.zip":          	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVh4ZlpzMzZsbDlObzJEaUUwLUtHUTBCTkN3c05Xeko1cDNpUlRVRDlDc1RIdz9lPTg1OGl6Mg==.zip",
    "./pkg/data/1000000to1999999_ERC20Transaction.zip":   	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRUU1ZFNmNGSkpGcWRjR3hJUDNJMkFCTW1GcXhTNmhVcjJYX08wV1JvdDZpdz9lPUd0czBRVQ==.zip",
    "./pkg/data/2000000to2999999_ERC20Transaction.zip":   	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWZtWUxJS1IyZWxOc2JmdGpmUnRGX2tCR2lLR0FhTDJkMWlnVUpfb2p5dnNmQT9lPVJrejhOYQ==.zip",
    "./pkg/data/3000000to3999999_ERC20Transaction.zip":   	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWM3R2hrbGZMV3RNcVZvdWs4ZFhOOEVCSEZNNE82YlBpekVTamNNY3J4UW9pZz9lPVF2QkVmOA==.zip",
    "./pkg/data/4000000to4999999_ERC20Transaction.zip":   	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWQ0REF2dEh5N3BLdEhtc2swaC1IdmtCLXZKTERPRHpIbGJ5S0pJSXRxNkRpdz9lPVF2cURNUQ==.zip",
    "./pkg/data/5000000to5999999_ERC20Transaction.zip":   	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWVZTE11MnBBOVZPbFkwTWFaU0hFMmNCZGxxZGZWUW1qQVlOOGljVEM0RC1adz9lPWpOcmUwbQ==.zip",
    "./pkg/data/6000000to6999999_ERC20Transaction.zip":   	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWJNdGM2NkRrcWhNakFhLUZsQnBndndCZVVtS2J6QS13WmNfOTQ4YjhocHBaUT9lPUZNRzA0bA==.zip",
    "./pkg/data/7000000to7999999_ERC20Transaction.zip":   	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWVWLVVrYTM0MFpNdjdtbmZ1dGZrZE1COGlXWEt1eHZoaHB1SThQSDJsRWx2dz9lPTlmSm1vcw==.zip",
    "./pkg/data/8000000to8999999_ERC20Transaction.zip":   	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVZsX0xWcGp5NTlEaFNzZUNSVkx2ZE1Cbm5KUnBXYlN1S1I4cDNfTG9Ba2VOQT9lPXlMWjA1Mw==.zip",
    "./pkg/data/9000000to9999999_ERC20Transaction.zip":   	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVpWVTd0NnpJUHhEa0MycWg3bHR6V0lCbFBYc1J1SEZkd2FRMU5BZ0RCMnFodz9lPTU4U1Vadw==.zip",
    "./pkg/data/10000000to10999999_ERC20Transaction.zip": 	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWZ4NGFWT09FXzlMb0U1am1rV0VLbElCMFpPWGh0ZU8tTFFiZG1ZR3NQX3dUUT9lPWJwRFFmdw==.zip",
    "./pkg/data/11000000to11999999_ERC20Transaction.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVdOTmY0NEY4cHhFcHU1RmhRQUJ5dVVCSnMtZThyY3B6dUlaWGpMdmxpdGhJUT9lPWJSMmZ3UQ==.zip",
	"./pkg/data/12000000to12999999_ERC20Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWVvcWRjYzJJdEZOakctbTdqVmpEdUlCOHhaRFFkWExUZDV4NWt2SjdsUGZ1dz9lPXdkUEw4OA.zip",
	"./pkg/data/13000000to13249999_ERC20Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVNuYkVHTmxDdk5OdG5kUU14bDZrV2dCa0JKeUFTX3JSUkE0d0RZMzVUTWtVdz9lPW5xaElxUA.zip",
	"./pkg/data/13250000to13499999_ERC20Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVJqNXNmTUVKcEJGanNXS1BPaldoS0FCZFZJUnpiRy1RRXgwMjNqX29Ub2VQZz9lPXg3OUVjQg.zip",
	"./pkg/data/13500000to13749999_ERC20Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVdGOGRhclBWN0JNa3NSWUNjVV9pekVCRVlUMmVPVlJ6enByVG5aZ0gxMlg4UT9lPWdiZFp1Yw.zip",
	"./pkg/data/13750000to13999999_ERC20Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWVKbXBFLXctUEZHcjRqbmFHTm1rUTBCS0NHNFNMc1pHeTFDRnFSa1FoczlIQT9lPWd3S2d0ZA.zip",
	"./pkg/data/14000000to14249999_ERC20Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVNVRDd0OTFiMXBBdXZIa2FYQlFQUUlCeTlNQXZrQUxNNmp4RXBhcVZhVUpxQT9lPWEzbWZleA.zip",
	"./pkg/data/14250000to14499999_ERC20Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVJ4S3VNMHFTQmRObkdjNG16Zm9zWXdCM19Rb19TQWU5cDZCQVk5YzJiak51QT9lPVN1TzRXaw.zip",
	"./pkg/data/stat_ERC20Transaction.py":					"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWFkUjdkWHlNRWRHZzdjdkt3TmpTMlFCOVVxdUJuSnh4eklSS1hKSzdoWXZDUT9lPWdFWlA0Qg==.py"
}
ERC721Transaction_Files = {
    "./pkg/data/0to999999_ERC721Transaction.zip":          "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVc5N213cGkwV0ZPdU4yLXhyUU1oclVCWlkxaDFNWWpOVGpQNUtaM2N3V0MyZz9lPVRKMnUycw==.zip",
    "./pkg/data/1000000to1999999_ERC721Transaction.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWNJdjE4NzZWaHRCajJmbHZXRVZ6VHdCcXRzdDFPR3k0U3NIa1Q5SzV2Yzk1dz9lPTZnYUt1cQ==.zip",
    "./pkg/data/2000000to2999999_ERC721Transaction.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVpYclB1QXF3S0pKa3hnNHJBNy1KZDBCV3RRc2hwMFZPN3I2R0x0S1I3aC0zUT9lPWwwcWU1cw==.zip",
    "./pkg/data/3000000to3999999_ERC721Transaction.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVRBRGVxdmZtZHhLdHRmTS0zR0NZMVFCenBIVTA0YVZqMV9oYmpRamdrWTlIUT9lPU1jQm1aVA==.zip",
    "./pkg/data/4000000to4999999_ERC721Transaction.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVhBQnNTcVFDNHhDaHJqcEVZdmFxSjhCMENISUVIc0ZqUGNka05aczJVeG9FZz9lPWliSmhCRA==.zip",
    "./pkg/data/5000000to5999999_ERC721Transaction.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWFUaWI4b3laaFZBcEs5R3ZXSmxvdVlCMEoxa3RlXzg0Mk1DQXF4a1VDcmRJZz9lPXFBMWtETw==.zip",
    "./pkg/data/6000000to6999999_ERC721Transaction.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVZtUFAyN3o4aXhEdkd3V2JDX2lzNklCYlQwOVVfR1pReFUtQWstZWFBOU5uQT9lPTFJUEcycw==.zip",
    "./pkg/data/7000000to7999999_ERC721Transaction.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVRscm1nTHp4REJKdUJINTIxd3VtdWtCaVJlOWtfMlZnZG9mWXk5UWRzNklNQT9lPWNIbm1haw==.zip",
    "./pkg/data/8000000to8999999_ERC721Transaction.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVlyZTJueU4yZ3RIdUxRdmx6MWlMZVVCbEQ4LWlQT3gxUHM2Rld5aUl3NUlUdz9lPU94bHB5Rw==.zip",
    "./pkg/data/9000000to9999999_ERC721Transaction.zip":   "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVY0T3F6VnBUVDlPZ0JYamNnMWMwNTBCUU9pRFBJQzd6bXN6cXRYa053akJhdz9lPVQ4T2hocw==.zip",
    "./pkg/data/10000000to10999999_ERC721Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRBTnpnQUI3cXhGbV9NVWF4NDBCYUlCSWN6YzE2M2VveG1od1daOHRCUngzUT9lPUpXVGc4ag==.zip",
    "./pkg/data/11000000to11999999_ERC721Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVJpb1U3UXI2STFNbDRQaFNNSG1jWDhCZllvMTBweHNkaVZWZHI2c09OVDREQT9lPWY3MU03eA==.zip",
	"./pkg/data/12000000to12999999_ERC721Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWVkVkxBMlFSSjVGbHdDSkttMmF5QWtCU09nWFZjN3hlUklzMlktWG5zVUFydz9lPWs1dVY4ZA.zip",
	"./pkg/data/13000000to13249999_ERC721Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRnWUwzUjc5SGxBdjE1TnZsZDVwUEVCcE04dXJ6ZnZOQjdUVnRvQ3hDVHhndz9lPWZyNGpCTg.zip",
	"./pkg/data/13250000to13499999_ERC721Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVg2V05PMUZuVDFPanVYczlWRE5HZEFCdGNCUTJBcmNyYUFYQ3A5aXh0YW4tUT9lPU1YZWdnUA.zip",
	"./pkg/data/13500000to13749999_ERC721Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVF3MWhKcUxLRmxBdW54OXRieXlfZDRCT2J4RFk0S2EyUGJpaEZXNDlXVV8xQT9lPUFlcE5YZg.zip",
	"./pkg/data/13750000to13999999_ERC721Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWNGc1IxOFVSRTVKbThIY01WMlU0czBCeEptZ2lVbUtBZlN3UGUzeFplWFpuZz9lPU03Z29IMA.zip",
	"./pkg/data/14000000to14249999_ERC721Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWVrbExkRHEyVkJDcHh0NU10ejY3UkFCOUVmTGxTNFQtQTZnYW1LVzBIdFQwZz9lPVRpN2pTaw.zip",
	"./pkg/data/14250000to14499999_ERC721Transaction.zip": "https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWZQc2F0TWlPdVpJdXBpSVZDY0tFc3dCam1nR2tMdUpFX09JSkdPTHhISm8xZz9lPW54VE1MMg.zip",
	"./pkg/data/stat_ERC721Transaction.py":				"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVY3LU56aldqQ3RPdVBKZmFveHVhWDRCRTNab2ZENm9LdVFJNzBVdDRfVnB5Zz9lPW5yZjNUcw==.py"
}
TokenInfo_Files = {
    "./pkg/data/ERC20TokenInfo.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRVZHV1pyWkRFek5Kdk5fblZVLW5yTk1CZjF5UVJJeUxUbF95ZzU4SjdTdVlnQT9lPTQzS3hCVA==.zip",
    "./pkg/data/ERC721TokenInfo.zip":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWNQVFNpT25FUXBJaDRGa0Q3d0QtTjRCZzhXb0xCVXBqMFRmY21ZU2loNXhJUT9lPXE2dDQzTg==.zip",
	"./pkg/data/stat_TokenInfo.py":	"https://link.jscdn.cn/sharepoint/aHR0cHM6Ly9vYmRvdGEtbXkuc2hhcmVwb2ludC5jbi86dTovZy9wZXJzb25hbC90Y3pwbF9vYmRvdGFfcGFydG5lcl9vbm1zY2hpbmFfY24vRWRoSnNxUUdyQnhCdmxSRVhwazgtVkVCWjdHeklySUpFUGlTeE1NcW1Lbks4UT9lPUgyZUlnYQ==.py"
}

All = [Block_Files, BlockTransaction_Files, InternalTransaction_Files, ContractInfo_Files, ERC20Transaction_Files, ERC721Transaction_Files, TokenInfo_Files]


def start():
	print("Select the datasets to download:")
	print("0. All")
	print("1. Block")
	print("2. Block Transaction")
	print("3. Internal Transaction ")
	print("4. Contract Info")
	print("5. ERC20 Transaction")
	print("6. ERC721 Transaction")
	print("7. Token Info")
	select = input("Please input a number (0~6): ")

	if select == "0":
		for Files in All:
			for localFile, srcUrl in Files.items():
				download(localFile, srcUrl)
	elif select == "1":
		for localFile, srcUrl in Block_Files.items():
			download(localFile, srcUrl)
	elif select == "2":
		for localFile, srcUrl in BlockTransaction_Files.items():
			download(localFile, srcUrl)
	elif select == "3":
		for localFile, srcUrl in InternalTransaction_Files.items():
			download(localFile, srcUrl)
	elif select == "4":
		for localFile, srcUrl in ContractInfo_Files.items():
			download(localFile, srcUrl)
	elif select == "5":
		for localFile, srcUrl in ERC20Transaction_Files.items():
			download(localFile, srcUrl)
	elif select == "6":
		for localFile, srcUrl in ERC721Transaction_Files.items():
			download(localFile, srcUrl)
	elif select == "7":
		for localFile, srcUrl in TokenInfo_Files.items():
			download(localFile, srcUrl)
	else:
		return start()

start()
print("finish")