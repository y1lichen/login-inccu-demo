import requests
from bs4 import BeautifulSoup

class GetHistory:
	def __init__(self, studentId, password):
		# for user data
		self.studentId = studentId
		self.password = password

		# for request
		self.session = requests.Session()
		self.session.headers['User-Agent'] = 'User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36'

		# for url
		self.loginUrl = 'https://i.nccu.edu.tw/Login.aspx?ReturnUrl=%2f'
		self.getSSOUrl = 'https://i.nccu.edu.tw/sso_app/PersonalInfoSSO.aspx?p=1&sid='+ self.studentId
		self.ssoLoginUrl = 'https://moltke.nccu.edu.tw/SSO/login.sso'
		# 全人網
		self.selfDevelopMenuUrl = 'https://moltke.nccu.edu.tw/selfDevelop_SSO/login.selfDevelop'
		self.targetUrl = 'https://moltke.nccu.edu.tw/selfDevelop_SSO/learningMapDataList.selfDevelop'

		self.loginPayLoad = {
			'__VIEWSTATE': '',
			'__VIEWSTATEGENERATOR': '',
			'__EVENTTARGET': "captcha$Login1$LoginButton",
			'__EVENTARGUMENT': '',
			'captcha$Login1$UserName': self.studentId,
			'captcha$Login1$Password': self.password,
			}

	def Login(self):
		loginHtml = self.session.get(self.loginUrl)
		loginParser = BeautifulSoup(loginHtml.text, 'html.parser')
		self.loginPayLoad['__VIEWSTATE'] = loginParser.select(
		"#__VIEWSTATE")[0]['value']
		self.loginPayLoad['__VIEWSTATEGENERATOR'] = loginParser.select(
		"#__VIEWSTATEGENERATOR")[0]['value']
		self.session.post(self.loginUrl, data=self.loginPayLoad)

	def getData(self):
		getssoHtml = self.session.get(self.getSSOUrl)
		getssoParse = BeautifulSoup(getssoHtml.text, 'html.parser')
		password = getssoParse.select("#password")[0]['value']
		data = {
			'id': self.studentId,
			'password': password,
			'p': 1,
			'url': self.selfDevelopMenuUrl
		}
		self.session.get(self.ssoLoginUrl, params=data)
		targetHtml = self.session.get(self.targetUrl)
		targetParse = BeautifulSoup(targetHtml.text, 'html.parser')
		# print(targetParse)

if __name__ == "__main__":
	test = GetHistory('<studentId>', '<password>')
	test.Login()
	test.getData()